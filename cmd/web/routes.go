package main

import (
	"github.com/bmizerany/pat"
	"github.com/ryantrue/web-application/pkg/config"
	"github.com/ryantrue/web-application/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
