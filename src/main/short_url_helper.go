package main

import (
	"log"
	"net/http"
	"time"
)

type ShortURLRequest struct {
	URL                 string        `json:"url"`
	CreateAt            string        `json:"created_at"`
	ExpirationInMinutes time.Duration `json:"expiration_in_minutes"`
}

func ParseReq(r *http.Request) ShortURLRequest {
	ua := r.UserAgent()
	log.Printf("Host: %s\n", r.Host)
	log.Printf("UserAgent: %s\n", ua)
	for name, value := range r.Header {
		log.Printf("[Header] %s: %s\n", name, value)
	}
	for _, cookie := range r.Cookies() {
		log.Printf("[Cookie] %s\n", cookie)
	}

	return ShortURLRequest{}
}
