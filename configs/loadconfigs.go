package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ConformityConfig struct {
	Resources []string
	Profiles  []string
    Db        DbConfig
    Rules     RulesConfig
}

type DbConfig struct {
    Type string
    Location string
}

type RulesConfig struct {
    RequiredTags []string `json:"required_tags"`
}

func LoadConfigs() *ConformityConfig {

	var config ConformityConfig
    setConfigDefaults(&config)

	configpath := "./conformitygopher.json"

	file, err := ioutil.ReadFile(configpath)
	if err != nil {
		log.Fatal("Error reading config", err)
	}

	json.Unmarshal(file, &config)
	log.Printf("Config Loaded: %s", configpath)
    log.Println(config)

	return &config
}

//Sets default values for config. Loading from Json will override these defaults
func setConfigDefaults(config *ConformityConfig) {
    config.Profiles = append(config.Profiles, "default")
    config.Resources = append(config.Resources, "ec2")
    config.Db.Type = "in-memory"
    config.Db.Location = "./conformitygopher.db"
}
