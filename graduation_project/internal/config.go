package internal

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	App struct {
		Port int     `yaml:"port"`
	} `yaml:"app"`
	Rpc struct {
		Port int     `yaml:"port"`
	} `yaml:"rpc"`
	Mysql string `yaml:"mysql"`
}

func AppConf() *Config {
	path := "./configs/app.yaml"

	var conf Config
	yamlFile, err := ioutil.ReadFile(path)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, &conf)
	}

	if err != nil {
		log.Fatalln("config parse error", err.Error())
	}

	return &conf
}