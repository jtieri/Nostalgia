package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	WebServer struct {
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		StaticFiles string `yaml:"static-files-path"`
	}
	MediaServer struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int16  `yaml:"port"`
		Name     string `yaml:"name"`
	}
}

func Load() *Config {
	data, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatal("Failed to read config file 'config.yml', check that it exists & try again.", err)
	}

	c := Config{}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatal("Failed to unmarshal config file 'config.yml', check that its format is correct & try again.", err)
	}

	log.Println("The file 'config.yml' has been successfully loaded.")
	return &c
}
