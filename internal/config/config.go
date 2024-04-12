package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/shynn12/biocad/pkg/logging"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type"`
		BindIp string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
	MongoDB struct {
		Host       string `json:"host"`
		Port       string `json:"port"`
		Database   string `json:"database"`
		Auth_db    string `json:"auth_db"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Collection string `json:"collection"`
	} `json:"mongodb"`
	Tsvpath string   `yaml:"tsvpath"`
	Pdfpath string   `yaml:"pdfpath"`
	Headers []string `yaml:"headers"`
	PerPage int      `yaml:"per_page"`
}

var instanse *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instanse = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instanse); err != nil {
			help, _ := cleanenv.GetDescription(instanse, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instanse
}
