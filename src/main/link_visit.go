package main

import (
	"log"
	"net/http"
	"time"
)

type LinkVisitedLog struct {
	Id         string    `json:"id"`
	RemoteAddr string    `json:"remote_addr"`
	ShortCode  string    `json:"short_code"`
	UA         string    `json:"ua"`
	Cookie     string    `json:"cookie"`
	VisitorId  string    `json:"visitor_id"`
	VisitedAt  time.Time `json:"visited_at"`
}

// GetVisitedLog 从http请求提取短链接访问记录
func GetVisitedLog(r *http.Request, shortCode string) *LinkVisitedLog {
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

// ParseShortUrlReq 解析访问记录
// 时间，地点，操作系统，浏览器，网络类型，
func ParseShortUrlReq(visitLog *LinkVisitedLog) {

}
