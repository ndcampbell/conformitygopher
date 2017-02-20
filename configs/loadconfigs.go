package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type BaseConfig struct {
	Resources []string
	Profiles  []string
	Db        DbConfig
	Rules     RulesConfig
}

type DbConfig struct {
	Type     string
	Location string
}

type RulesConfig struct {
	RequiredTags []string `json:"required_tags"`
}

var Config BaseConfig

func LoadConfigs() {

	setConfigDefaults(&Config)

	configpath := "./conformitygopher.json"

	file, err := ioutil.ReadFile(configpath)
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}

	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal("Error with config json: ", err)
	}
	log.Printf("Config Loaded: %s", configpath)
}

//Sets default values for config. Loading from Json will override these defaults
func setConfigDefaults(config *BaseConfig) {
	config.Profiles = append(config.Profiles, "default")
	config.Resources = append(config.Resources, "ec2")
	config.Db.Type = "in-memory"
	config.Db.Location = "./conformitygopher.db"
}
