package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonConfig struct {
	Resources []string
	Profiles  []string
}

func LoadConfigs() {
	var config jsonConfig

	file, err := ioutil.ReadFile("./conformitygopher.json")
	if err != nil {
		fmt.Println("Error reading config", err)
		os.Exit(1)
	}

	json.Unmarshal(file, &config)
	fmt.Printf("%v\n", config)
}
