package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

const SiteUrl = "iiu8.cn/"

func (app *App) createShortURL(w http.ResponseWriter, r *http.Request) {
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

	s, err := app.Config.S.Shorten(req.URL, req.ExpirationInMinutes)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusCreated, shortenResp{ShortURL: SiteUrl + s, Code: 2000, Message: "OK"})
	}
}

func (app *App) getShortURLInfo(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	s := values.Get("short_url")

	// fmt.Printf("get info: %s\n", s)
	data, err := app.Config.S.ShortURLInfo(s)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusOK, data)
	}
}

func (app *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	url, err := app.Config.S.UnShorten(vars["short_url"])
	if err != nil {
		respondWithError(w, err)
	} else {
		// TODO parse req and save it
		ParseReq(r)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func respondWithError(w http.ResponseWriter, err error) {
	log.Println(err)
	switch e := err.(type) {
	case Error:
		log.Printf("HTTP %d - %s", e.Status(), e)
		respondWithJSON(w, e.Status(), e.Error())
	default:
		// TODO need json, with code, message ...
		respondWithJSON(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
