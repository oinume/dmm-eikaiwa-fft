package mux

import (
	"github.com/oinume/lekcije/server/web/api"
	"github.com/oinume/lekcije/server/web/middleware"
	"goji.io"
	"goji.io/pat"
	"github.com/oinume/lekcije/server/web"
)

func Create() *goji.Mux {
	mux := goji.NewMux()
	mux.UseC(middleware.AccessLogger)
	mux.UseC(middleware.SetDbToContext)
	mux.UseC(middleware.SetLoggedInUserToContext)
	mux.UseC(middleware.LoginRequiredFilter)

	mux.HandleFuncC(pat.Get("/"), web.Index)
	mux.HandleFuncC(pat.Get("/logout"), web.Logout)
	mux.HandleFuncC(pat.Get("/oauth/google"), web.OAuthGoogle)
	mux.HandleFuncC(pat.Get("/oauth/google/callback"), web.OAuthGoogleCallback)
	mux.HandleFuncC(pat.Post("/me/followingTeachers/create"), web.PostMeFollowingTeachersCreate)
	mux.HandleFuncC(pat.Post("/me/followingTeachers/delete"), web.PostMeFollowingTeachersDelete)

	mux.HandleFuncC(pat.Get("/api/status"), api.GetStatus)
	mux.HandleFuncC(pat.Get("/api/me/followingTeachers"), api.GetMeFollowingTeachers)
	return mux
}
