package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"github.com/mananwalia959/bookings-app/pkg/config"
	"github.com/mananwalia959/bookings-app/pkg/handlers"
)

var app *config.AppConfig

func GetRoutes(a *config.AppConfig) http.Handler {
	app = a
	mux := chi.NewRouter()
	// mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(csrfMiddleware)
	mux.Use(sessionMiddleWare)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}

func csrfMiddleware(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.UseTls,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func sessionMiddleWare(next http.Handler) http.Handler {
	return app.SessionManager.LoadAndSave(next)
}
