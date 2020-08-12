package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// App encapsulates Env, Router and middleware
type App struct {
	Router      *mux.Router
	Middlewares *Middleware
	Config      *Env
}

// Initialize is initialization of app
func (app *App) Initialize(env *Env) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app.Config = env
	app.Router = mux.NewRouter()
	app.Middlewares = &Middleware{}

	app.initializeRoutes()
}

func (app *App) initializeRoutes() {
	m := alice.New(app.Middlewares.LoggingHandler,
		app.Middlewares.RecoverHandler,
		app.Middlewares.CorsHeadersHandler)
	app.Router.Handle("/api/shorten", m.ThenFunc(app.createLink)).Methods(http.MethodPost, http.MethodOptions)
	app.Router.Handle("/api/info", m.ThenFunc(app.getLinkInfo)).Methods(http.MethodGet, http.MethodOptions)
	app.Router.Use(mux.CORSMethodMiddleware(app.Router))

	app.Router.Handle("/{url:[a-zA-Z0-9]{1,11}}", m.ThenFunc(app.redirect)).Methods(http.MethodGet)
}

// Run the app
func (app *App) Run(addr string) {
	log.Printf("Server running on: %s", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
