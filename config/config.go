package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Mysql struct {
		DataSources string `yaml:"dataSource"`
	} `yaml:"mysql"`
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
	JWT struct {
		Salt        string `yaml:"salt"`
		Expiredtime int    `yaml:"expiredtime"`
	} `yaml:"jwt"`
	Mail struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mail"`
}

var once sync.Once
var Conf *Config = new(Config)

func init() {
	once.Do(func() {
		wd := os.Getenv("WORKSPACE_DIR")
		if wd != "" {
			os.Chdir(wd)
		}
		yamlFile, err := os.ReadFile("config/config.yaml")
		if err != nil {
			log.Fatalf("yamlfile.Get err   #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, Conf)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	})
}
