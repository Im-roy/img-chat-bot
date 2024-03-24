package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Host        string `yaml:"host" envconfig:"DB_HOST"`
	Port        int    `yaml:"port" envconfig:"DB_PORT"`
	User        string `yaml:"user" envconfig:"DB_USER"`
	Password    string `yaml:"password" envconfig:"DB_PASSWORD"`
	Name        string `yaml:"name" envconfig:"DB_NAME"`
	Timeout     int    `yaml:"timeout" envconfig:"DB_CONNECTION_TIMEOUT"`
	MaxOpenConn int    `yaml:"maxOpenConn" envconfig:"DB_MAX_OPEN_CONNECTIONS"`
	MaxIdleConn int    `yaml:"maxIdleConn" envconfig:"DB_MAX_IDLE_CONNECTIONS"`
}

type HttpConfig struct {
	Address string `yaml:"address" envconfig:"LISTEN_ADDRESS"`
}

type AppConfig struct {
	DbConfig   DbConfig   `yaml:"db"`
	HttpConfig HttpConfig `yaml:"http"`
}

var (
	config     *AppConfig
	configPath string
)

func loadConfig(filepath string) error {
	yamlData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlData, &config)
}

func GetConfig(path string) (AppConfig, error) {
	if config == nil {
		err := loadConfig(path)
		if err != nil {
			return *config, err
		}
	}
	return *config, nil
}
