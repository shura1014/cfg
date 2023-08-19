package cfg

import (
	"github.com/shura1014/cfg/g"
	"github.com/shura1014/common/utils/stringutil"
	"github.com/shura1014/common/utils/timeutil"
	"strconv"
	"strings"
	"time"
)

type localAdapter struct {
	data map[string]any
	envs map[string]string
}

func (c *localAdapter) Get(pattern string) (value any, err error) {
	var (
		keys = strings.Split(pattern, Point)
		data any
	)
	// 先从环境变量取
	if v, ok := c.envs[pattern]; ok {
		return v, nil
	}
	data = c.data
	for index, key := range keys {
		if d := checkAndGet(key, data); d != nil {
			if index == len(keys)-1 {
				return *d, nil
			} else {
				data = *d
			}
		} else {
			break
		}
	}
	return nil, nil
}

func (c *localAdapter) GetString(key string, def ...string) string {
	v, _ := c.Get(key)
	if v != nil {
		return stringutil.ToString(v)
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (c *localAdapter) GetBool(key string, def ...bool) bool {
	v, _ := c.Get(key)
	if v != nil {
		return v.(bool)
	}
	if len(def) > 0 {
		return def[0]
	}
	return false
}

func (c *localAdapter) GetStringMap(key string) map[string]any {
	v, _ := c.Get(key)
	if v != nil {
		return v.(map[string]any)
	}
	return nil
}

func (c *localAdapter) GetArray(key string) (value []any) {
	v, _ := c.Get(key)
	if v != nil {
		return v.([]any)
	}
	return nil
}

func (c *localAdapter) GetInt64(key string, def ...int64) int64 {
	v, _ := c.Get(key)
	if v != nil {
		return v.(int64)
	}

	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (c *localAdapter) GetInt(key string, def ...int) int {
	v := c.GetString(key)
	if v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			g.Error(err)
			return 0
		}
		return i
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (c *localAdapter) GetTime(key string, def ...time.Duration) (value time.Duration) {
	v := c.GetString(key)
	if v != "" {
		return timeutil.Convert(v)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (c *localAdapter) GetAll() (data map[string]any, err error) {
	return c.data, nil
}

func (c *localAdapter) PutAll(data map[string]any) {
	for k, v := range data {
		c.data[k] = v
	}
}

func checkAndGet(key string, data any) *any {
	switch value := data.(type) {
	case map[string]any:
		if v, ok := value[key]; ok {
			return &v
		}
	case []any:
		if stringutil.IsNumeric(key) {
			n, err := strconv.Atoi(key)
			if err == nil && len(value) > n {
				return &value[n]
			}
		}
	}
	return nil
}
