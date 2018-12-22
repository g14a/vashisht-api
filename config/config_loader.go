package config

import (
	"io/ioutil"
	"log"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

var appConfigInstance *AppConfig
var once sync.Once

// GetAppConfig returns an application configuration
func GetAppConfig() *AppConfig {
	once.Do(func() {
		loadConfig()
	})

	return appConfigInstance
}

func loadConfig() {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &appConfigInstance)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
