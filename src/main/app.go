package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"gopkg.in/validator.v2"
)

// App encapsulates Env, Router and middleware
type App struct {
	Router      *mux.Router
	Middlewares *Middleware
	Config      *Env
}

type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"min=0"`
}

type shortlinkResp struct {
	Shortlink string `json:"shortlink"`
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
	m := alice.New(app.Middlewares.LoggingHandler, app.Middlewares.RecoverHabdler)
	app.Router.Handle("/api/shorten", m.ThenFunc(app.createShortlink)).Methods("POST")
	app.Router.Handle("/api/info", m.ThenFunc(app.getShortlinkInfo)).Methods("GET")
	app.Router.Handle("/{shortlink:[a-zA-Z0-9]{1,11}}", m.ThenFunc(app.redirect)).Methods("GET")
}

func (app *App) createShortlink(w http.ResponseWriter, r *http.Request) {
	var req shortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, StatusError{http.StatusBadRequest, fmt.Errorf("parse parameters failed %v", r.Body)})
		return
	}
	if err := validator.Validate(req); err != nil {
		respondWithError(w, StatusError{http.StatusBadRequest, fmt.Errorf("validate parameters failed %v", req)})
		return
	}
	defer r.Body.Close()

	// fmt.Printf("create: %v\n", req)
	s, err := app.Config.S.Shorten(req.URL, req.ExpirationInMinutes)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusCreated, shortlinkResp{Shortlink: s})
	}
}

func (app *App) getShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	s := vals.Get("shortlink")

	// fmt.Printf("get info: %s\n", s)
	data, err := app.Config.S.ShortlinkInfo(s)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusOK, data)
	}
}

func (app *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// fmt.Printf("short link: %s\n", vars["shortlink"])
	unShortlink, err := app.Config.S.UnShorten(vars["shortlink"])
	if err != nil {
		respondWithError(w, err)
	} else {
		http.Redirect(w, r, unShortlink, http.StatusTemporaryRedirect)
	}
}

// Run the app
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func respondWithError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case Error:
		log.Printf("HTTP %d - %s", e.Status(), e)
		respondWithJSON(w, e.Status(), e.Error())
	default:
		respondWithJSON(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
