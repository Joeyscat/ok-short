package main

type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes uint32 `json:"expiration_in_minutes" validate:"min=0,max=1440"`
}

type shortenResp struct {
	Link    string `json:"link"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
