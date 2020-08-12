package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	. "github.com/joeyscat/ok-short/common"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

func (app *App) createLink(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	var req shortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, StatusError{Code: http.StatusBadRequest, Err: fmt.Errorf("parse parameters failed %v", r.Body)})
		return
	}
	if err := validator.Validate(req); err != nil {
		respondWithError(w, StatusError{Code: http.StatusBadRequest, Err: fmt.Errorf("validate parameters failed %v", req)})
		return
	}
	defer r.Body.Close()

	s, err := app.Config.API.Shorten(req.URL, req.ExpirationInMinutes)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusCreated, shortenResp{Link: LinkPrefix + s, Code: 2000, Message: "OK"})
	}
}

func (app *App) getLinkInfo(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	s := values.Get("short_url")

	// fmt.Printf("get info: %s\n", s)
	link, err := app.Config.API.LinkInfo(s)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusOK, link)
	}
}

func (app *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sc := vars["url"]
	url, err := app.Config.API.UnShorten(sc)
	if err != nil {
		respondWithError(w, err)
	} else {
		// TODO 加入队列异步处理
		l := getVisitedLog(r, sc)
		app.Config.API.StoreVisitedLog(l)
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

func getVisitedLog(r *http.Request, shortCode string) *LinkVisitedLog {
	reqLog := LinkVisitedLog{
		RemoteAddr: r.Header.Get("Remote_addr"),
		ShortCode:  shortCode,
		UA:         r.UserAgent(),
		Cookie:     r.Header.Get("Cookie"),
		VisitorId:  "0",
		VisitedAt:  Now(),
	}
	log.Printf("ReqLog: %+v\n", reqLog)
	return &reqLog
}

// --------------------------- admin ----------------------------

func (app *App) login(w http.ResponseWriter, r *http.Request) {

	app.Config.ADMIN.Login("", "")
}
