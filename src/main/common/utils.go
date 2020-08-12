package common

import (
	"github.com/mattheath/base62"
	"time"
)

const encodeCus = "4HY6KuGCXIyvWQLUnJMF2AST7whfgOxo3NDeiRdBr51kpcPjtl9bqEzVaZs80m"

var customEncoding = base62.NewEncoding(encodeCus)

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
