package common

import (
	"github.com/Unknwon/goconfig"
	"sync"
	"log"
)

type ConfigInter interface {
	GetValue(string) string
	MustValue(string) string
}


type Config struct {
	cfg *goconfig.ConfigFile
}

var config *Config
var once sync.Once

func NewConfig() *Config{
	once.Do(func() {
		config = new(Config)
		configFile := ProjectConfigPath + "/config.ini"
		cfg, err := goconfig.LoadConfigFile(configFile)
		if err != nil {
			log.Fatalf("无法加载配置文件：%s", err)
		}
		config.cfg = cfg
	})

	return config
}


func (c *Config) GetValue(defaultVal string, key string) string {
	value, err := c.cfg.GetValue("", key)
	if err != nil {
		return defaultVal
	}
	return value
}

func (c *Config) GetSectionValue(section string, key string, defaultVal string) string {
	value := c.cfg.MustValue(section, key, defaultVal)
	return value
}