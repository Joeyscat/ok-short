package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	. "github.com/joeyscat/ok-short/common"
	"github.com/joeyscat/ok-short/store"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"strconv"
)

func (app *App) createLink(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	var req ShortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, StatusError{Code: ParamPostBodyInvalid,
			Err: fmt.Errorf("%s %v", BSText(ParamPostBodyInvalid), err.Error())})
		return
	}
	if err := validator.Validate(req); err != nil {
		respondWithError(w, http.StatusBadRequest, StatusError{Code: ParamIllegal,
			Err: fmt.Errorf("%s %v", BSText(ParamIllegal), err.Error())})
		return
	}
	defer r.Body.Close()

	s, err := app.Config.API.Shorten(req.URL, req.ExpirationInMinutes)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	} else {
		respondWithJSON(w, http.StatusCreated, Resp{
			Code:    Success,
			Message: BSText(Success),
			Data:    ShortenRespData{Link: store.LinkPrefix + s},
		})
	}
}

func (app *App) getLinkInfo(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	s := values.Get("short_url")
	if s == "" {
		respondWithError(w, http.StatusBadRequest, StatusError{Code: ParamURLInvalid,
			Err: fmt.Errorf(BSText(ParamURLInvalid))})
	}

	// fmt.Printf("get info: %s\n", s)
	link, err := app.Config.API.LinkInfo(s)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	} else {
		link.Id = 0
		respondWithJSON(w, http.StatusOK, Resp{
			Code:    Success,
			Message: BSText(Success),
			Data:    LinkInfoRespData{Link: *link},
		})
	}
}

func (app *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sc := vars["url"]
	url, err := app.Config.API.UnShorten(sc)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	} else {
		// TODO 加入队列异步处理
		l := getVisitedLog(r, sc)
		app.Config.API.StoreVisitedLog(l)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

// 返回错误响应，status为HTTP状态码, 业务代码默认为系统错误
// 当err参数为自定义的Error时，从中提取code作为业务代码
func respondWithError(w http.ResponseWriter, status int, err error) {
	switch e := err.(type) {
	case Error:
		log.Printf("HTTP %d - %s", status, e)
		respondWithJSON(w, status, Resp{
			Code:    e.ECode(),
			Message: e.Error(),
		})
	default:
		respondWithJSON(w, status, Resp{
			Code:    SystemErr,
			Message: err.Error(),
		})
	}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload Resp) {
	resp, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
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

func (app *App) register(w http.ResponseWriter, r *http.Request) {
	var req RegisterReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, StatusError{Code: ParamPostBodyInvalid,
			Err: fmt.Errorf("%s %v", BSText(ParamPostBodyInvalid), err.Error())})
		return
	}
	if err := validator.Validate(req); err != nil {
		respondWithError(w, http.StatusBadRequest, StatusError{Code: ParamIllegal,
			Err: fmt.Errorf("%s %v", BSText(ParamIllegal), err.Error())})
		return
	}
	defer r.Body.Close()
	success, err := app.Config.ADMIN.Register(req.Name, req.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	} else {
		if !success {
			respondWithJSON(w, http.StatusOK, Resp{Code: UserRegisterFail, Message: BSText(UserRegisterFail)})
		} else {
			respondWithJSON(w, http.StatusOK, Resp{Code: Success, Message: BSText(Success)})
		}
	}
}

func (app *App) login(w http.ResponseWriter, r *http.Request) {
	var req LoginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, StatusError{Code: ParamPostBodyInvalid,
			Err: fmt.Errorf("%s %v", BSText(ParamPostBodyInvalid), err.Error())})
		return
	}
	if err := validator.Validate(req); err != nil {
		respondWithError(w, http.StatusBadRequest, StatusError{Code: ParamIllegal,
			Err: fmt.Errorf("%s %v", BSText(ParamIllegal), err.Error())})
		return
	}
	defer r.Body.Close()
	token, err := app.Config.ADMIN.Login(req.Name, req.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	} else {
		respondWithJSON(w, http.StatusOK, Resp{Code: Success, Message: BSText(Success), Data: LoginRespData{Token: token}})
	}
}

func (app *App) adminInfo(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	userInfoStr, err := app.Config.ADMIN.UserInfo(token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	var user AdminInfoRespData
	_ = json.Unmarshal([]byte(userInfoStr), &user)
	//userInfo, _ := json.Unmarshal(userInfoStr)
	respondWithJSON(w, http.StatusOK, Resp{Code: Success, Message: BSText(Success), Data: user})
}

func (app *App) links(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	page, _ := strconv.ParseInt(values.Get("page"), 10, 32)
	size, _ := strconv.ParseInt(values.Get("size"), 10, 32)

	links, count, total, err := app.Config.ADMIN.QueryLinks(uint32(page), uint32(size))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, Resp{
		Code:    Success,
		Message: BSText(Success),
		Data:    QueryLinksRespData{Total: total, Count: count, Links: links},
	})
}
