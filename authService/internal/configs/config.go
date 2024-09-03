package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Server struct {
	Timeout time.Duration `yaml:"timeout"`
	Port    string        `yaml:"port"`
}

type Config struct {
	Server `yaml:"server"`
}

var config Config

func Get() *Config {
	return &config
}

func init() {
	file, err := ioutil.ReadFile("./internal/configs/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	log.Println("Config is loaded")
}
