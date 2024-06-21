package config

import (
	"encoding/json"
	"github.com/run-bigpig/xpvideo/internal/types"
	"github.com/run-bigpig/xpvideo/internal/utils"
	"os"
	"slices"
	"sync"
)

var (
	once       sync.Once
	globalConf *Config
)

type Config struct {
	Source        []*types.Source
	DefaultSource *types.Source
}

func Set() {
	once.Do(func() {
		globalConf = &Config{}
		makeDefaultSource()
		parseJson("./source.json")
		for _, v := range globalConf.Source {
			SetDefaultSource(v.Key)
			break
		}
	})
}

func SetDefaultSource(source string) {
	for _, v := range globalConf.Source {
		if v.Key == source {
			globalConf.DefaultSource = v
			break
		}
	}
}

func Get() *Config {
	return globalConf
}

func parseJson(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	source := make([]*types.Source, 0)
	err = json.Unmarshal(data, &source)
	if err != nil {
		panic(err)
	}
	keyMap := make(map[string]struct{})
	//移除有相同key的source
	for index, v := range source {
		if _, ok := keyMap[v.Key]; ok {
			source = slices.Delete(source, index, index+1)
			continue
		}
		keyMap[v.Key] = struct{}{}
	}
	globalConf.Source = source
}

// 初始化默认配置
func makeDefaultSource() {
	if !utils.FileExists("./source.json") {
		defaultSource := []*types.Source{{
			Name: "360资源站",
			Key:  "360zy",
			Url:  "https://360zy.com/api.php/provide/vod",
		}}
		data, err := json.Marshal(defaultSource)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile("./source.json", data, 0644)
	}
}
