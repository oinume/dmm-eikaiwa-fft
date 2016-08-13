package web

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/oinume/lekcije/server/util"
	"github.com/pkg/errors"
)

const ApiTokenCookieName = "apiToken"

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func ListenPort() int {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4001"
	}
	p, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		return -1
	}
	return int(p)
}

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

func InternalServerError(w http.ResponseWriter, err error) {
	//switch _ := errors.Cause(err).(type) { // TODO:
	//default:
	// unknown error
	http.Error(w, fmt.Sprintf("Internal Server Error\n\n%v", err), http.StatusInternalServerError)
	if e, ok := err.(stackTracer); ok {
		for _, f := range e.StackTrace() {
			fmt.Fprintf(w, "%+v\n", f)
		}
	}
	//}
}
