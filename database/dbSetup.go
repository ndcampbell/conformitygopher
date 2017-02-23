package database

import (
	"log"

	"github.com/ndcampbell/conformitygopher/configs"
)

func DbSetup(config *configs.DbConfig) {
	if config.Type == "in-memory" {
		db := BoltStruct{}
		db.BoltSetup(config.Location)
		log.Println("Bolt DB Setup")
	}

}
