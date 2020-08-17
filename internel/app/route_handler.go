package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joeyscat/ok-short/internel/pkg"
	. "github.com/joeyscat/ok-short/internel/pkg/common"
	"github.com/joeyscat/ok-short/internel/pkg/model"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (app *App) createLink(w http.ResponseWriter, r *http.Request) {
	var req ShortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusOK, StatusError{Code: ParamPostBodyInvalid,
			Err: fmt.Errorf("%s %v", BSText(ParamPostBodyInvalid), err.Error())})
		return
	}
	if err := validator.Validate(req); err != nil {
		respondWithError(w, http.StatusOK, StatusError{Code: ParamIllegal,
			Err: fmt.Errorf("%s %v", BSText(ParamIllegal), err.Error())})
		return
	}
	defer r.Body.Close()

	s, err := app.Context.API.Shorten(req.URL, req.ExpirationInMinutes)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	} else {
		respondWithJSON(w, http.StatusCreated, Resp{
			Code:    Success,
			Message: BSText(Success),
			Data:    ShortenRespData{Link: pkg.LinkPrefix + s},
		})
	}
}

func (app *App) getLinkInfo(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	sc := values.Get("sc")
	if sc == "" {
		respondWithError(w, http.StatusOK, StatusError{Code: ParamURLInvalid,
			Err: fmt.Errorf(BSText(ParamURLInvalid))})
	}

	link, err := app.Context.API.LinkInfo(sc)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	} else {
		respondWithJSON(w, http.StatusOK, Resp{Code: Success, Message: BSText(Success), Data: link})
	}
}

func (app *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sc := vars["url"]
	url, err := app.Context.API.UnShorten(sc)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
	} else {
		// TODO 加入队列异步处理
		l := getVisitedLog(r)
		app.Context.API.StoreVisitedLog(l)
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

func getVisitedLog(r *http.Request) *model.LinkTrace {
	reqLog := model.LinkTrace{
		Sid:    Sid(),
		Ip:     r.Header.Get("Remote_addr"),
		URL:    getURL(r),
		UA:     r.UserAgent(),
		Cookie: r.Header.Get("Cookie"),
	}
	log.Printf("ReqLog: %+v\n", reqLog)
	return &reqLog
}

func getURL(r *http.Request) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	return strings.Join([]string{scheme, r.Host, r.RequestURI}, "")
}

// --------------------------- admin ----------------------------

func (app *App) register(w http.ResponseWriter, r *http.Request) {
	var req RegisterReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusOK, StatusError{Code: ParamPostBodyInvalid,
			Err: fmt.Errorf("%s %v", BSText(ParamPostBodyInvalid), err.Error())})
		return
	}
	if req.Name == "" || req.Password == "" {
		respondWithError(w, http.StatusOK, StatusError{Code: ParamPasswordEmpty,
			Err: errors.New(BSText(ParamAccOrPassEmpty))})
		return
	}
	if err := validator.Validate(req); err != nil {
		respondWithError(w, http.StatusOK, StatusError{Code: ParamIllegal,
			Err: fmt.Errorf("%s %v", BSText(ParamIllegal), err.Error())})
		return
	}
	defer r.Body.Close()
	success, err := app.Context.ADMIN.Register(req.Name, req.Password)
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
		respondWithError(w, http.StatusOK, StatusError{Code: ParamPostBodyInvalid,
			Err: fmt.Errorf("%s %v", BSText(ParamPostBodyInvalid), err.Error())})
		return
	}
	if req.Name == "" || req.Password == "" {
		respondWithError(w, http.StatusOK, StatusError{Code: ParamPasswordEmpty,
			Err: errors.New(BSText(ParamAccOrPassEmpty))})
		return
	}
	if err := validator.Validate(req); err != nil {
		respondWithError(w, http.StatusOK, StatusError{Code: ParamIllegal,
			Err: fmt.Errorf("%s %v", BSText(ParamIllegal), err.Error())})
		return
	}
	defer r.Body.Close()
	token, err := app.Context.ADMIN.Login(req.Name, req.Password)
	if err != nil {
		respondWithError(w, http.StatusOK, err)
	} else {
		respondWithJSON(w, http.StatusOK, Resp{Code: Success, Message: BSText(Success), Data: LoginRespData{Token: token}})
	}
}

func (app *App) adminUser(w http.ResponseWriter, r *http.Request) {
	//token := r.Header.Get("Authorization")
	token := r.URL.Query().Get("token")
	if token == "" {
		respondWithError(w, http.StatusOK, StatusError{Code: ParamTokenEmpty,
			Err: errors.New(BSText(ParamTokenEmpty))})
		return
	}
	userInfoStr, err := app.Context.ADMIN.UserInfo(token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	var user AdminInfoRespData
	_ = json.Unmarshal([]byte(userInfoStr), &user)
	// TODO hard code
	user.Roles = append(user.Roles, "admin")
	user.AvatarURL = "https://avatars3.githubusercontent.com/u/27766600?s=460&u=ac9809d85b4986bb38b85c1ec79bbebec476b574&v=4"
	//userInfo, _ := json.Unmarshal(userInfoStr)
	respondWithJSON(w, http.StatusOK, Resp{Code: Success, Message: BSText(Success), Data: user})
}

func (app *App) linkList(w http.ResponseWriter, r *http.Request) {
	page, limit := getPageParams(r)

	links, totalCount, err := app.Context.ADMIN.QueryLinkList(page, limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	itemCount := len(*links)
	respondWithJSON(w, http.StatusOK, Resp{
		Code:    Success,
		Message: BSText(Success),
		Data:    QueryListRespData{TotalCount: totalCount, ItemCount: uint32(itemCount), Item: &links},
	})
}

func (app *App) linkTraceList(w http.ResponseWriter, r *http.Request) {
	page, limit := getPageParams(r)

	traceList, totalCount, err := app.Context.ADMIN.QueryLinkTraceList(page, limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	itemCount := len(*traceList)
	respondWithJSON(w, http.StatusOK, Resp{
		Code:    Success,
		Message: BSText(Success),
		Data:    QueryListRespData{TotalCount: totalCount, ItemCount: uint32(itemCount), Item: &traceList},
	})
}

func (app *App) adminUserList(w http.ResponseWriter, r *http.Request) {
	page, limit := getPageParams(r)

	admins, totalCount, err := app.Context.ADMIN.QueryAdminUserList(page, limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	itemCount := len(*admins)
	respondWithJSON(w, http.StatusOK, Resp{
		Code:    Success,
		Message: BSText(Success),
		Data:    QueryListRespData{TotalCount: totalCount, ItemCount: uint32(itemCount), Item: &admins},
	})
}

func getPageParams(r *http.Request) (uint32, uint32) {
	values := r.URL.Query()
	page, _ := strconv.ParseInt(values.Get("page"), 10, 32)
	limit, _ := strconv.ParseInt(values.Get("limit"), 10, 32)
	// page>=1,20>=limit>=1
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 20 {
		limit = 20
	}
	return uint32(page), uint32(limit)
}
