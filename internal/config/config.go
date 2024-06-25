package config

import (
	"encoding/json"
	"github.com/run-bigpig/xpvideo/internal/constant"
	"github.com/run-bigpig/xpvideo/internal/types"
	"github.com/run-bigpig/xpvideo/internal/utils"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"sync"
)

const (
	ConfigFile = "config.yaml"
)

var (
	once sync.Once
)

type SourceConfig struct {
	List []*types.Source `yaml:"list"`
}

type Config struct {
	Source *SourceConfig `yaml:"source"`
}

func Init() {
	once.Do(func() {
		viper.SetConfigFile(ConfigFile)
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		if !utils.FileExists("./" + ConfigFile) {
			makeDefaultConfig()
		}
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
	})
}

// 生成默认配置
func makeDefaultConfig() {
	c := &Config{}
	c.Source = &SourceConfig{}
	list, err := getHttpSourceList()
	if err != nil {
		panic(err)
	}
	c.Source.List = list
	viper.SetDefault("source", c.Source)
	err = viper.WriteConfig()
	if err != nil {
		panic(err)
	}
}

// GetDefaultSourceUrl 获取默认源
func GetDefaultSourceUrl() string {
	var sourceList []*types.Source
	err := viper.UnmarshalKey("source.list", &sourceList)
	if err != nil {
		return ""
	}
	for _, source := range sourceList {
		if source.Default {
			return source.Url
		}
	}
	if len(sourceList) > 0 {
		return sourceList[0].Url
	}
	list, err := getHttpSourceList()
	if err == nil {
		return list[0].Url
	}
	return ""
}

func getHttpSourceList() ([]*types.Source, error) {
	var list []*types.Source
	//获取默认源
	resp, err := http.Get(constant.DefaultUrl)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
