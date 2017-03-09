package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/jpillora/go-ogle-analytics"
	"github.com/oinume/lekcije/server/config"
	"github.com/oinume/lekcije/server/context_data"
	"github.com/oinume/lekcije/server/controller/flash_message"
	"github.com/oinume/lekcije/server/errors"
	"github.com/oinume/lekcije/server/logger"
	"github.com/oinume/lekcije/server/util"
	"github.com/stvp/rollbar"
	"github.com/uber-go/zap"
)

const (
	APITokenCookieName   = "apiToken"
	TrackingIDCookieName = "trackingId"
)

func TemplateDir() string {
	if util.IsProductionEnv() {
		return "static"
	} else {
		return "src/html"
	}
}

func TemplatePath(file string) string {
	return path.Join(TemplateDir(), file)
}

func ParseHTMLTemplates(files ...string) *template.Template {
	f := []string{
		TemplatePath("_base.html"),
		TemplatePath("_flashMessage.html"),
	}
	f = append(f, files...)
	return template.Must(template.ParseFiles(f...))
}

func InternalServerError(w http.ResponseWriter, err error) {
	//switch _ := errors.Cause(err).(type) { // TODO:
	//default:
	// unknown error
	if rollbar.Token != "" {
		rollbar.Error(rollbar.ERR, err)
	}
	fields := []zap.Field{
		zap.Error(err),
	}
	if e, ok := err.(errors.StackTracer); ok {
		b := &bytes.Buffer{}
		for _, f := range e.StackTrace() {
			fmt.Fprintf(b, "%+v\n", f)
		}
		fields = append(fields, zap.String("stacktrace", b.String()))
	}
	logger.App.Error("InternalServerError", fields...)

	http.Error(w, fmt.Sprintf("Internal Server Error\n\n%v", err), http.StatusInternalServerError)
	if !config.IsProductionEnv() {
		fmt.Fprintf(w, "----- stacktrace -----\n")
		if e, ok := err.(errors.StackTracer); ok {
			for _, f := range e.StackTrace() {
				fmt.Fprintf(w, "%+v\n", f)
			}
		}
	}
}

func JSON(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, `{ "status": "Failed to Encode as JSON" }`, http.StatusInternalServerError)
		return
	}
}

type commonTemplateData struct {
	StaticURL         string
	GoogleAnalyticsID string
	CurrentURL        string
	CanonicalURL      string
	TrackingID        string
	IsUserAgentPC     bool
	IsUserAgentSP     bool
	IsUserAgentTablet bool
	UserID            string
	NavigationItems   []navigationItem
	FlashMessage      *flash_message.FlashMessage
}

type navigationItem struct {
	Text string
	URL  string
}

var loggedInNavigationItems = []navigationItem{
	{"ホーム", "/me"},
	{"設定", "/me/setting"},
	{"ログアウト", "/me/logout"},
}

var loggedOutNavigationItems = []navigationItem{
	{"ホーム", "/"},
}

func getCommonTemplateData(req *http.Request, loggedIn bool, userID uint32) commonTemplateData {
	canonicalURL := fmt.Sprintf("%s://%s%s", config.WebURLScheme(req), req.Host, req.RequestURI)
	canonicalURL = (strings.SplitN(canonicalURL, "?", 2))[0] // TODO: use url.Parse
	data := commonTemplateData{
		StaticURL:         config.StaticURL(),
		GoogleAnalyticsID: config.GoogleAnalyticsID(),
		CurrentURL:        req.RequestURI,
		CanonicalURL:      canonicalURL,
		IsUserAgentPC:     util.IsUserAgentPC(req),
		IsUserAgentSP:     util.IsUserAgentSP(req),
		IsUserAgentTablet: util.IsUserAgentTablet(req),
	}

	if loggedIn {
		data.NavigationItems = loggedInNavigationItems
	} else {
		data.NavigationItems = loggedOutNavigationItems
	}
	if flashMessageKey := req.FormValue("flashMessageKey"); flashMessageKey != "" {
		flashMessage, _ := flash_message.MustStore(req.Context()).Load(flashMessageKey)
		data.FlashMessage = flashMessage
	}
	data.TrackingID = context_data.MustTrackingID(req.Context())
	if userID != 0 {
		data.UserID = fmt.Sprint(userID)
	}

	return data
}

var gaHTTPClient *http.Client = &http.Client{
	Transport: &logger.LoggingHTTPTransport{DumpHeaderBody: true},
	Timeout:   time.Second * 7,
}

const (
	eventCategoryUser             = "user"
	eventCategoryFollowingTeacher = "followingTeacher"
)

func sendMeasurementEvent(req *http.Request, category, action, label string, value int64, userID uint32) {
	gaClient, err := ga.NewClient(os.Getenv("GOOGLE_ANALYTICS_ID"))
	if err != nil {
		logger.App.Warn("ga.NewClient() failed", zap.Error(err))
	}
	gaClient.HttpClient = gaHTTPClient
	gaClient.UserAgentOverride(req.UserAgent())

	gaClient.ClientID(context_data.MustTrackingID(req.Context()))
	gaClient.DocumentHostName(req.Host)
	gaClient.DocumentPath(req.URL.Path)
	gaClient.DocumentTitle(req.URL.Path)
	gaClient.DocumentReferrer(req.Referer())
	gaClient.IPOverride(GetRemoteAddress(req))

	logFields := []zap.Field{
		zap.String("category", category),
		zap.String("action", action),
	}
	event := ga.NewEvent(category, action)
	if label != "" {
		event.Label(label)
		logFields = append(logFields, zap.String("label", label))
	}
	if value != 0 {
		event.Value(value)
		logFields = append(logFields, zap.Int64("value", value))
	}
	if userID != 0 {
		gaClient.UserID(fmt.Sprint(userID))
		logFields = append(logFields, zap.Uint("userID", uint(userID)))
	}
	if err := gaClient.Send(event); err == nil {
		// TODO: stats log
		logger.App.Debug("sendMeasurementEvent() success", logFields...)
	} else {
		logger.App.Warn("gaClient.Send() failed", zap.Error(err))
	}
}

func GetRemoteAddress(req *http.Request) string {
	xForwardedFor := req.Header.Get("X-Forwarded-For")
	if xForwardedFor == "" {
		return (strings.Split(req.RemoteAddr, ":"))[0]
	}
	return strings.TrimSpace((strings.Split(xForwardedFor, ","))[0])
}
