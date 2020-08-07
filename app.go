package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
)

// App encapsulates Env, Router and middleware
type App struct {
	Router *mux.Router
}

type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"min=0"`
}

type shortlinkResp struct {
	Shortlink string `json:"shortlink"`
}

// Initialize is initialization of app
func (app *App) Initialize() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/api/shorten", app.createShortlink).Methods("POST")
	app.Router.HandleFunc("/api/info", app.getShortlinkInfo).Methods("GET")
	app.Router.HandleFunc("/{shortlink:[a-zA-Z0-9]{1,11}}", app.redirect).Methods("GET")
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

	fmt.Printf("create: %v\n", req)
}

func (app *App) getShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	s := vals.Get("shortlink")
	fmt.Printf("get info: %s\n", s)
}

func (app *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("short link: %s\n", vars["shortlink"])
}

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
