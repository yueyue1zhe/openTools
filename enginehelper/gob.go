package core

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"io"
)

func GobEncodeAny[T any](p T) (res []byte, err error) {
	js, err := json.Marshal(&p)
	if err != nil {
		return
	}
	useStr := base64.StdEncoding.EncodeToString(js)
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err = enc.Encode(useStr); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 接收 gob 编码的字节切片并返回解码后的结构体
func GobDecodeAny[T any](data io.Reader) (res T, err error) {
	var raw string
	dec := gob.NewDecoder(data)
	err = dec.Decode(&raw)
	if err != nil {
		return
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return
	}
	err = json.Unmarshal(decodeBytes, &res)
	return
}
