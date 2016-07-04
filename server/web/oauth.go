package web

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	google_auth2 "google.golang.org/api/oauth2/v2"

	"github.com/oinume/lekcije/server/model"
)

var googleOAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Endpoint:     google.Endpoint,
	RedirectURL:  fmt.Sprintf("http://localhost:%d/oauth/google/callback", 4000),
	Scopes: []string{
		"openid email",
		"openid profile",
	},
}

func OAuthGoogle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	state := randomString(32)
	cookie := &http.Cookie{
		Name:     "oauthState",
		Value:    state,
		Path:     "/",
		Expires:  time.Now().Add(time.Minute * 30),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, googleOAuthConfig.AuthCodeURL(state), http.StatusFound)
}

func OAuthGoogleCallback(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if err := checkState(r); err != nil {
		internalServerError(w, err)
		return
	}
	token, idToken, err := exchange(r)
	if err != nil {
		internalServerError(w, err)
		return
	}
	name, email, err := getNameAndEmail(token, idToken)
	if err != nil {
		internalServerError(w, err)
		return
	}
	db, err := model.Open()
	if err != nil {
		internalServerError(w, errors.Wrap(err, "Failed to connect db: %v"))
		return
	}

	user := model.User{Name: name, Email: email}
	if err := db.FirstOrCreate(&user, model.User{Email: email}).Error; err != nil {
		internalServerError(w, errors.Wrap(err, "Failed to access user"))
		return
	}

	data := map[string]interface{}{
		"id":          user.Id,
		"name":        user.Name,
		"email":       user.Email,
		"accessToken": token.AccessToken,
		"idToken":     idToken,
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		internalServerError(w, errors.Errorf("Failed to encode JSON"))
		return
	}
}

func randomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func checkState(r *http.Request) error {
	state := r.FormValue("state")
	oauthState, err := r.Cookie("oauthState")
	if err != nil {
		return errors.Wrap(err, "Failed to get cookie oauthState")
	}
	if state != oauthState.Value {
		return errors.Wrap(err, "state mismatch")
	}
	return nil
}

func exchange(r *http.Request) (*oauth2.Token, string, error) {
	code := r.FormValue("code")
	token, err := googleOAuthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, "", errors.Wrap(err, "Failed to exchange")
	}
	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, "", errors.Errorf("Failed to get id_token")
	}
	return token, idToken, nil
}

func getNameAndEmail(token *oauth2.Token, idToken string) (string, string, error) {
	oauth2Client := oauth2.NewClient(oauth2.NoContext, oauth2.StaticTokenSource(token))
	service, err := google_auth2.New(oauth2Client)
	if err != nil {
		return "", "", errors.Wrap(err, "Failed to create oauth2.Client")
	}

	userinfo, err := service.Userinfo.V2.Me.Get().Do()
	if err != nil {
		return "", "", errors.Wrap(err, "Failed to get userinfo")
	}

	tokeninfo, err := service.Tokeninfo().IdToken(idToken).Do()
	if err != nil {
		return "", "", errors.Wrap(err, "Failed to get tokeninfo")
	}

	return userinfo.Name, tokeninfo.Email, nil
}
