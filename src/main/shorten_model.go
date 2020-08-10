package main

type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"min=0"`
}

type shortenResp struct {
	ShortURL string `json:"shortURL"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
}
