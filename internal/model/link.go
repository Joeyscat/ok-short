package model

type Link struct {
	Sc        string `bson:"sc" json:"sc"`                 // 短链代码
	Status    string `bson:"status" json:"status"`         // 短链状态
	OriginURL string `bson:"origin_url" json:"origin_url"` // 原始链接
	Exp       uint32 `bson:"exp" json:"exp"`               // 过期时间
}
