package app

import (
	"github.com/joeyscat/ok-short/internel/pkg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// App encapsulates Context, Router and middleware
type App struct {
	Router     *mux.Router
	Middleware *pkg.Middleware
	Context    *Context
}

// Initialize is initialization of app
func (app *App) Initialize(c *Context) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app.Context = c
	app.Router = mux.NewRouter()
	app.Middleware = &pkg.Middleware{}

	app.initializeRoutes()
}

func (app *App) initializeRoutes() {
	m := alice.New(app.Middleware.LoggingHandler,
		app.Middleware.RecoverHandler,
		app.Middleware.CorsHandler,
	)
	app.Router.Handle("/api/shorten", m.ThenFunc(app.createLink)).Methods(http.MethodPost, http.MethodOptions)
	app.Router.Handle("/api/info", m.ThenFunc(app.getLinkInfo)).Methods(http.MethodGet, http.MethodOptions)

	app.Router.Handle("/admin-api/register", m.ThenFunc(app.register)).Methods(http.MethodPost, http.MethodOptions)
	app.Router.Handle("/admin-api/user/login", m.ThenFunc(app.login)).Methods(http.MethodPost, http.MethodOptions)
	app.Router.Handle("/admin-api/user/info", m.ThenFunc(app.adminUser)).Methods(http.MethodGet, http.MethodOptions)
	app.Router.Handle("/admin-api/user/list", m.ThenFunc(app.adminUserList)).Methods(http.MethodGet, http.MethodOptions)
	app.Router.Handle("/admin-api/link/list", m.ThenFunc(app.linkList)).Methods(http.MethodGet, http.MethodOptions)
	app.Router.Handle("/admin-api/link/trace/list", m.ThenFunc(app.linkTraceList)).Methods(http.MethodGet, http.MethodOptions)

	app.Router.Handle("/{url:[a-zA-Z0-9]{1,11}}", m.ThenFunc(app.redirect)).Methods(http.MethodGet)

	app.Router.Use(mux.CORSMethodMiddleware(app.Router))
}

// Run the app
func (app *App) Run(addr string) {
	log.Printf("Server running on: %s", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
