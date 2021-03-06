package emailer

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/oinume/lekcije/backend/logger"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type transport struct {
	called bool
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.called = true
	resp := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusOK,
		Status:     "OK",
		Body:       ioutil.NopCloser(strings.NewReader("OK")),
	}
	return resp, nil
}

func TestSendGridSender_Send(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	s := `
From: lekcije@lekcije.com
To: gmail <oinume@gmail.com>, oinume@lampetty.net
Subject: テスト
Body: text/html
oinume さん
こんにちは
	`
	template := NewTemplate("TestNewEmailFromTemplate", strings.TrimSpace(s))
	data := struct {
		Name  string
		Email string
	}{
		"oinume",
		"oinume@gmail.com",
	}
	email, err := NewEmailFromTemplate(template, data)
	r.Nil(err)

	email.SetCustomArg("userId", "1")
	email.SetCustomArg("teacherIds", "1,2,3")
	tr := &transport{}
	err = NewSendGridSender(
		&http.Client{Transport: tr},
		logger.NewAppLogger(os.Stdout, zapcore.InfoLevel),
	).Send(context.Background(), email)
	r.Nil(err)
	a.True(tr.called)
}
