package cfg

import (
	"fmt"
	"github.com/shura1014/cfg/g"
	"github.com/shura1014/cfg/goproperties"
	"github.com/shura1014/cfg/gotoml"
	"github.com/shura1014/cfg/goyaml"
	"github.com/shura1014/common/env"
	"github.com/shura1014/common/goerr"
	"github.com/shura1014/common/utils/fileutil"
	"github.com/shura1014/common/utils/stringutil"
	"os"
	"runtime"
	"strings"
)

var (
	DefaultConfigName  = "app.toml"
	DefaultConfigDir   = "config"
	ConfigNameEnv      = "app.config.filename"
	ConfigNameDir      = "app.config.dir"
	supportedFileTypes = []string{"toml", "yaml", "yml", "xml", "properties"}
)

const (
	Point = "."
)

// LoadConfig 每个不同的中间件应当初始化自己的配置文件
func LoadConfig(dir string, fileNames ...string) (conf *Config, err error) {
	dir = deduceDir(dir)
	// 	加载配置文件数据
	configs, err := GetConfigFilePath(dir, fileNames...)
	if err != nil {
		return nil, err
	}

	conf = NewLocal()
	for _, configPath := range configs {
		ext := fileutil.ExtName(configPath)
		// 判断是否是可支持的文件
		if stringutil.IsArray(supportedFileTypes, ext) {
			bytes := fileutil.ReadBytes(configPath)
			ReadToConf(conf, ext, configPath, bytes)
		}
	}

	// 环境变量优先
	envs := env.GetAll()
	adapter := conf.ConfigAdapter.(*localAdapter)
	adapter.envs = envs
	return
}

func LoadEnv() (conf *Config, err error) {
	envs := env.GetAll()
	conf = NewLocal()
	adapter := conf.ConfigAdapter.(*localAdapter)
	adapter.envs = envs
	return
}

type Config struct {
	ConfigAdapter
}

func ReadToConf(conf *Config, ext string, fileName string, content []byte) {
	var (
		data    = make(map[string]any)
		err     error
		adapter = conf.ConfigAdapter.(*localAdapter)
	)
	switch ext {
	case "toml":
		err = gotoml.DecodeTo(content, &data)
	case "yaml", "yml":
		err = goyaml.DecodeTo(content, &data)
	case "properties":
		data, err = goproperties.Decode(content)
	}

	if err != nil {
		panic(goerr.Text("%s decode failed", fileName))
	}

	adapter.PutAll(data)
}

func GetConfigFilePath(dir string, fileNames ...string) ([]string, error) {
	var (
		fileList []string
		ok       = false
		envFile  = ""
	)
	// 传入的dir为空，那么尝试着环境变量获取一下，环境变量获取为空，那么使用默认的dir
	if dir == "" {
		if dir, ok = env.GetEnv(ConfigNameDir); !ok {
			dir = fileutil.JoinCurrent(DefaultConfigDir)
		}
	}
	dir = fileutil.RealPath(dir)
	s := os.Args
	fmt.Println(s)
	if len(fileNames) <= 0 || fileNames[0] == "" {
		if envFile, ok = env.GetEnv(ConfigNameEnv); !ok {
			envFile = DefaultConfigName
		}
		fileNames = strings.Split(envFile, ",")
	}

	foundFiles := fileutil.ListFileName(dir)
	for _, foundName := range foundFiles {
		for _, fileName := range fileNames {
			if strings.HasPrefix(foundName, fileName) {
				fileList = append(fileList, fileutil.Join(dir, fileName))
			}
		}
	}

	if len(fileList) <= 0 {
		text := goerr.Text("not found config")
		g.Error(text)
		return nil, text
	}

	return fileList, nil
}

func deduceDir(dir string) string {
	if dir == "" {
		return ""
	}
	if strings.HasPrefix(dir, "/") {
		// 绝对路径
		return dir
	}

	// 	相对路径
	var realPath string
	_, file, _, _ := runtime.Caller(2)
	fileutil.DirFunc(fileutil.Dir(file), func() {
		realPath = fileutil.RealPath(dir)
	})
	return realPath
}
