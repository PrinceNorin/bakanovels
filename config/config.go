package config

import (
	"sync"

	"github.com/jinzhu/configor"
)

var once sync.Once
var config *Config

var configPaths = []string{"config/config.yml", "config/default_config.yml"}

func Get() *Config {
	once.Do(func() {
		config = &Config{}
	})
	return config
}

func init() {
	newConfig := &Config{
		Host:        Get().Host,
		Port:        Get().Port,
		Environment: Get().Environment,
	}
	config = newConfig
	configor.Load(config, configPaths...)
}
