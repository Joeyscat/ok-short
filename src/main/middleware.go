package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// 用于接受Token的Header
const TokenHeaderName = "OK-Short-Token"

type Middleware struct {
}

// LoggingHandler 记录http请求日志
func (m Middleware) LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//header := r.Header
		//remoteAddr := header.Get("Remote_addr")
		//realIp := header.Get("X-Real-Ip")
		//forwardedFor := header.Get("X-Forwarded-For")
		//log.Printf("[Header] Remote_addr:%s | X-Real-Ip:%s | X-Forwarded-For:%s\n", remoteAddr, realIp, forwardedFor)

		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

// RecoverHandler recover panic
func (m Middleware) RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic: %+v\n", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// CorsHandler 添加跨域请求必须的Header，并直接放行Options请求
func (m Middleware) CorsHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Index(r.URL.String(), "/api/") == 0 ||
			strings.Index(r.URL.String(), "/admin-api/") == 0 {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length, Authorization, cross-request-open-sign, "+TokenHeaderName)
			w.Header().Set("Access-Control-Max-Age", "1728000")
		}
		if r.Method == http.MethodOptions {
			// handle preflight in here
		} else {
			next.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(fn)
}
