package common

type ShortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes uint32 `json:"expiration_in_minutes" validate:"min=0,max=1440,nonnil"`
}

type ShortenRespData struct {
	Link string `json:"link"`
}

type LinkRespData struct {
	Sid    string `json:"sid"` // 业务标识
	URL    string `json:"url"` // 短链
	Status string `json:"status"`
	//Group     Group     `json:"group"` // 分组
	Name      string `json:"name"`
	OriginURL string `json:"origin_url"` // 原始链接
	//PV     PV     `json:"pv"`
	CreatedAt string `json:"created_at"`
}

type QueryListRespData struct {
	TotalCount uint32      `json:"total_count"`
	ItemCount  uint32      `json:"item_count"`
	Item       interface{} `json:"item"`
}

type LinkTraceRespData struct {
	Sid       string `json:"sid"`
	URL       string `json:"url"`
	UA        string `json:"ua"`
	Ip        string `json:"ip"`
	CreatedAt string `json:"visited_at"`
}

type RegisterReq struct {
	Name     string `json:"name" validate:"min=5,max=20"`
	Password string `json:"password" validate:"min=64,max=64"`
}

type LoginReq struct {
	Name     string `json:"name" validate:"min=5,max=20"`
	Password string `json:"password" validate:"min=64,max=64"`
}

type LoginRespData struct {
	Token string `json:"token"`
}

type AdminInfoRespData struct {
	Name      string   `json:"name"`
	Roles     []string `json:"roles"`
	AvatarURL string   `json:"avatar"`
}

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
