// Package gotoml 提供TOML格式内容解析为结构体对象，结构体对象反解析
// go get github.com/BurntSushi/toml
package gotoml

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/shura1014/cfg/g"
	"github.com/shura1014/common/goerr"
	"reflect"
)

func Decode(v []byte) (any, error) {
	var result any
	if err := toml.Unmarshal(v, &result); err != nil {
		g.Error(err)
		err = goerr.Wrapf(err, "toml decode failed")
		return nil, err
	}
	return result, nil
}

func DecodeTo(v []byte, result any) (err error) {

	t := reflect.TypeOf(result)
	if t.Kind() != reflect.Pointer {
		return goerr.Text("the acceptor must be pointer type")
	}
	err = toml.Unmarshal(v, result)
	if err != nil {
		g.Error(err)
		err = goerr.Wrapf(err, "toml decode failed")
	}
	return err
}

func Encode(v any) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(buffer).Encode(v); err != nil {
		g.Error(err)
		err = goerr.Wrapf(err, "toml encode failed")
		return nil, err
	}
	return buffer.Bytes(), nil
}
