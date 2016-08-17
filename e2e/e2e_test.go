package e2e

import (
	"bytes"
	"testing"
	"fmt"
	"os"
	"net/http/httptest"
	"net/http"
	"net"
	"time"

	"github.com/oinume/lekcije/server/web"
	"github.com/oinume/lekcije/server/mux"
	"github.com/oinume/lekcije/server/logger"
	"github.com/uber-go/zap"
)

var server *httptest.Server
var client = http.DefaultClient

func TestMain(m *testing.M) {
	var accessLogBuffer bytes.Buffer
	var appLogBuffer bytes.Buffer
	logger.AccessLogger = zap.New(zap.NewJSONEncoder(), zap.Output(zap.AddSync(&accessLogBuffer)))
	logger.AppLogger = zap.New(zap.NewJSONEncoder(), zap.Output(zap.AddSync(&appLogBuffer)))
	port := web.ListenPort()
	mux := mux.Create()
	port += 1
	server = newTestServer(mux, port)
	fmt.Printf("Test HTTP server created: port=%d, url=%s\n", port, server.URL)
	defer server.Close()

	client.Timeout = 5 * time.Second
	os.Chdir("../")
	status := m.Run()
	defer os.Exit(status)
}

// newTestServer returns a new test Server with fixed port number.
func newTestServer(handler http.Handler, port int) *httptest.Server {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		if listener, err = net.Listen("tcp6", fmt.Sprintf("[::1]:%d", port)); err != nil {
			panic(fmt.Sprintf("httptest: failed to listen on a port: %v", err))
		}
	}
	ts := &httptest.Server{
		Listener: listener,
		Config:   &http.Server{Handler: handler},
	}
	ts.Start()
	return ts
}