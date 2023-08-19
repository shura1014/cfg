package goproperties

import (
	"bytes"
	"github.com/magiconair/properties"
	"github.com/shura1014/cfg/g"
	"github.com/shura1014/common/goerr"
	"github.com/shura1014/common/utils/stringutil"
	"strings"
)

func Decode(content []byte) (data map[string]any, err error) {
	data = make(map[string]any)
	prop, err := properties.Load(content, properties.UTF8)
	if err != nil || prop == nil {
		g.Error(err)
		return nil, goerr.Wrapf(err, "properties load failed")
	}

	for _, key := range prop.Keys() {
		value, _ := prop.Get(key)
		paths := strings.Split(key, ".")
		propToMap(data, paths, value)
	}
	return data, nil
}

func Encode(data map[string]any) (res []byte, err error) {
	prop := properties.NewProperties()
	prop.Sort()
	receive := map[string]any{}
	mapToProp(receive, data, "", ".")
	for key, value := range receive {
		_, _, err = prop.Set(key, stringutil.ToString(value))
		if err != nil {
			g.Error(err)
			err = goerr.Wrapf(err, "properties set key:%s value:%v failed", key, value)
			return nil, err
		}
	}

	// 排个序
	prop.Sort()

	var buf bytes.Buffer

	_, err = prop.Write(&buf, properties.UTF8)
	if err != nil {
		err = goerr.Wrapf(err, "Properties Write buf failed")
		return nil, err
	}

	return buf.Bytes(), nil
}

// propToMap 统一格式转换
// a.b.c=1
// a[b[c]]=1
func propToMap(data map[string]any, paths []string, value string) {

	for index, path := range paths {
		if index == len(paths)-1 {
			data[path] = value
			return
		}
		v, ok := data[path]
		// 没找到
		if !ok {
			// 创建一个
			m := make(map[string]any)
			data[path] = m
			data = m
			continue
		}
		// 找到了，强转
		m, ok := v.(map[string]any)
		if !ok {
			m = make(map[string]any)
			data[path] = m
		}
		data = m
	}
}

// mapToProp 统一格式转换
// a[b[c]]=1
// a.b.c=1
func mapToProp(receive map[string]any, data map[string]any, prefix string, delimiter string) {
	if receive != nil && prefix != "" && receive[prefix] != nil {
		return
	}
	var m map[string]any
	if prefix != "" {
		prefix += delimiter
	}

	for k, val := range data {
		fullKey := prefix + k
		switch val.(type) {
		case map[string]any:
			m = val.(map[string]any)
		default:
			receive[strings.ToLower(fullKey)] = val
			continue
		}
		mapToProp(receive, m, fullKey, delimiter)
	}
}
