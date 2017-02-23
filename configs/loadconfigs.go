package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type BaseConfig struct {
	DbConfig
	RulesConfig
	Resources []string
	Profiles  []string
}

type DbConfig struct {
	Type     string
	Location string
}

type RulesConfig struct {
	RequiredTags []string `json:"required_tags"`
}

func LoadConfigs() {

	config := &BaseConfig{}
	config.SetConfigDefaults()

	configpath := "./conformitygopher.json"

	file, err := ioutil.ReadFile(configpath)
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}

	err = json.Unmarshal(file, config)
	if err != nil {
		log.Fatal("Error with config json: ", err)
	}
	log.Printf("Config Loaded: %s", configpath)
}

//Sets default values for config. Loading from Json will override these defaults
func (b *BaseConfig) SetConfigDefaults() {
	b.Profiles = []string{}
	b.Resources = []string{}
	b.Profiles = append(b.Profiles, "default")
	b.Resources = append(b.Resources, "ec2")
	b.Type = "in-memory"
	b.Location = "./conformitygopher.db"
}
