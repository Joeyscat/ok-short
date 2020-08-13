package main

import . "github.com/joeyscat/ok-short/common"

type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes uint32 `json:"expiration_in_minutes" validate:"min=0,max=1440"`
}

type shortenResp struct {
	Link    string `json:"link"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type queryLinksReq struct {
	Page uint32 `json:"page" validate:"min=0"`
	Size uint32 `json:"size" validate:"min=0,max=20"`
}

type queryLinksResp struct {
	Total uint32  `json:"total"`
	Count uint32  `json:"count"`
	Links *[]Link `json:"links"`
}
