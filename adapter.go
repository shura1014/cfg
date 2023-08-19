package cfg

import "time"

type ConfigAdapter interface {
	Get(pattern string) (value any, err error)

	GetString(pattern string, def ...string) (value string)

	GetInt64(pattern string, def ...int64) (value int64)

	GetInt(pattern string, def ...int) (value int)

	GetBool(pattern string, def ...bool) (value bool)

	GetTime(pattern string, def ...time.Duration) (value time.Duration)

	GetArray(pattern string) (value []any)

	GetStringMap(pattern string) (value map[string]any)

	GetAll() (data map[string]any, err error)
}
