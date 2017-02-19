package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ConformityConfig struct {
	Resources []string
	Profiles  []string
}

func LoadConfigs() *ConformityConfig {

	var config ConformityConfig
	configpath := "./conformitygopher.json"

	file, err := ioutil.ReadFile(configpath)
	if err != nil {
		log.Fatal("Error reading config", err)
	}

	json.Unmarshal(file, &config)
	log.Printf("Config Loaded: %s", configpath)

	return &config
}
