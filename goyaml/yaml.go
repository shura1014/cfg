package goyaml

import (
	"bytes"
	"github.com/shura1014/cfg/g"
	"github.com/shura1014/common/goerr"
	"gopkg.in/yaml.v3"
	"strings"
)

func Decode(content []byte) (map[string]any, error) {
	data := make(map[string]any)
	err := yaml.Unmarshal(content, &data)
	if err != nil {
		g.Error(err)
		return nil, goerr.Wrapf(err, "yaml parse exception")
	}
	return data, nil
}

func DecodeTo(content []byte, data any) error {
	err := yaml.Unmarshal(content, data)
	if err != nil {
		g.Error(err)
		return goerr.Wrapf(err, "yaml parse exception")
	}
	return nil
}

func Encode(value any) (out []byte, err error) {
	if out, err = yaml.Marshal(value); err != nil {
		g.Error(err)
		err = goerr.Wrapf(err, "yaml encode failed")
	}
	return
}

func EncodeIndent(value interface{}, indent string) (out []byte, err error) {
	out, err = Encode(value)
	if err != nil {
		return
	}
	if indent != "" {
		var (
			buffer = bytes.NewBuffer(nil)
			array  = strings.Split(strings.TrimSpace(string(out)), "\n")
		)
		for _, v := range array {
			buffer.WriteString(indent)
			buffer.WriteString(v)
			buffer.WriteString("\n")
		}
		out = buffer.Bytes()
	}
	return
}
