package common

import (
	"github.com/lithammer/shortuuid/v3"
	"github.com/mattheath/base62"
	"github.com/teris-io/shortid"
	"time"
)

const encodeCus = "4HY6KuGCXIyvWQLUnJMF2AST7whfgOxo3NDeiRdBr51kpcPjtl9bqEzVaZs80m"

var customEncoding = base62.NewEncoding(encodeCus)
var ShortId = shortid.Shortid{}

func init() {
	ShortId = *newShortId()
}

func newShortId() *shortid.Shortid {
	s, err := shortid.New(1, shortid.DefaultABC, 9527)
	if err != nil {
		panic(err)
	}
	return s
}

func Sid() string {
	generate, _ := ShortId.Generate()
	return generate
}

// Base62Encode Custom base62 encoding
func Base62Encode(i int64) string {
	return customEncoding.EncodeInt64(i)
}

func Now() time.Time {
	//location, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	return time.Now()
	//}
	//return time.Now().In(location)
	return time.Now()
}

func ShortUUID() string {
	return shortuuid.NewWithNamespace("http://example.com")
}
