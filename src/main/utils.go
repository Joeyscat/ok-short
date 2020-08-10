package main

import "github.com/mattheath/base62"

const encodeCus = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var customEncoding = base62.NewEncoding(encodeCus)

// Base62Encode Custom base62 encoding
func Base62Encode(i int64) string {
	return customEncoding.EncodeInt64(i)
}
