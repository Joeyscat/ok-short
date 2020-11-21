package codec

import (
	"bytes"
	"encoding/gob"
)

func Encoder(s interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(s)
	return b.Bytes(), err
}

func Decoder(b []byte, s interface{}) error {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	return dec.Decode(s)
}
