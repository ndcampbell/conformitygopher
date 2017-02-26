package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
)

type BaseConfig struct {
	Db        DbConfig
	Rules     RulesConfig
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

func LoadConfigs() *BaseConfig {

	config := &BaseConfig{}
	config.SetConfigDefaults()

	usr, _ := user.Current()
	homedir := usr.HomeDir
	configpath := homedir + "/conformitygopher.json"
	log.Println(configpath)

	file, err := ioutil.ReadFile(configpath)
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}

	err = json.Unmarshal(file, config)
	if err != nil {
		log.Fatal("Error with config json: ", err)
	}
	log.Printf("Config Loaded: %s", configpath)

	return config
}

//Sets default values for config. Loading from Json will override these defaults
func (b *BaseConfig) SetConfigDefaults() {
	b.Profiles = append(b.Profiles, "default")
	b.Resources = append(b.Resources, "ec2")
	b.Db.Type = "in-memory"
	b.Db.Location = "./conformitygopher.db"
}
