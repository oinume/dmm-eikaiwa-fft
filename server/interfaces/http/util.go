package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"

	"github.com/oinume/lekcije/server/config"
	"github.com/oinume/lekcije/server/context_data"
	"github.com/oinume/lekcije/server/errors"
	"github.com/oinume/lekcije/server/interfaces/http/flash_message"
	"github.com/oinume/lekcije/server/logger"
	"github.com/oinume/lekcije/server/util"
	"github.com/stvp/rollbar"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	APITokenCookieName   = "apiToken"
	TrackingIDCookieName = "trackingId"
)

func TemplateDir() string {
	return "frontend/html"
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
	fields := []zapcore.Field{
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
	Text      string
	URL       string
	NewWindow bool
}

var loggedInNavigationItems = []navigationItem{
	{"ホーム", "/me", false},
	{"設定", "/me/setting", false},
	{"お問い合わせ", "https://goo.gl/forms/CIGO3kpiQCGjtFD42", true},
	{"ログアウト", "/me/logout", false},
}

var loggedOutNavigationItems = []navigationItem{
	{"ホーム", "/", false},
}

func getCommonTemplateData(req *http.Request, loggedIn bool, userID uint32) commonTemplateData {
	canonicalURL := fmt.Sprintf("%s://%s%s", config.WebURLScheme(req), req.Host, req.RequestURI)
	canonicalURL = (strings.SplitN(canonicalURL, "?", 2))[0] // TODO: use url.Parse
	data := commonTemplateData{
		StaticURL:         config.StaticURL(),
		GoogleAnalyticsID: config.DefaultVars.GoogleAnalyticsID,
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

func GetRemoteAddress(req *http.Request) string {
	xForwardedFor := req.Header.Get("X-Forwarded-For")
	if xForwardedFor == "" {
		return (strings.Split(req.RemoteAddr, ":"))[0]
	}
	return strings.TrimSpace((strings.Split(xForwardedFor, ","))[0])
}