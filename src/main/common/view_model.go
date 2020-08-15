package common

type ShortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes uint32 `json:"expiration_in_minutes" validate:"min=0,max=1440,nonnil"`
}

type ShortenRespData struct {
	Link string `json:"link"`
}

type LinkInfoRespData struct {
	Link Link `json:"link"`
}

type QueryLinksRespData struct {
	Total uint32  `json:"total"`
	Count uint32  `json:"count"`
	Links *[]Link `json:"links"`
}

type RegisterReq struct {
	Name     string `json:"name" validate:"min=6,max=20"`
	Password string `json:"password" validate:"min=64,max=64"`
}

type LoginReq struct {
	Name     string `json:"name" validate:"min=6,max=20"`
	Password string `json:"password" validate:"min=64,max=64"`
}

type LoginRespData struct {
	Token string `json:"token"`
}

type AdminInfoRespData struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrResp(code int, msg string) Resp {
	return Resp{
		Code:    code,
		Message: msg,
	}
}

func OKResp(code int, msg string, data interface{}) Resp {
	return Resp{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
