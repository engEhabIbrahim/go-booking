package main

import (
	"github.com/engEhabIbrahim/go-booking/pkg/config"
	"github.com/engEhabIbrahim/go-booking/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app config.AppConfig) http.Handler {

	// pat routing package
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// chi routing package
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer) // recover from panic ( built in )
	// mux.Use(WriteToConsole)       // custom middleware
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
