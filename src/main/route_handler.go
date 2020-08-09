package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

func (app *App) createShortlink(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
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
	s, err := app.Config.S.Shorten(req.URL, req.ExpirationInMinutes)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusCreated, shortlinkResp{Shortlink: s, Code: 2000, Message: "OK"})
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
