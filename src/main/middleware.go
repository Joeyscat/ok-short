package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

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

// CorsHeadersHandler add headers for CORS Request
func (m Middleware) CorsHeadersHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Index(r.URL.String(), "/api") == 0 {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length, Authorization, Accept,X-Requested-With")
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
