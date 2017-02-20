package database

import (
	"log"

	"github.com/ndcampbell/conformitygopher/configs"
)

func DbSetup() {
	if configs.Config.Db.Type == "in-memory" {
		BoltSetup(configs.Config.Db.Location)
		log.Println("Bolt DB Setup")
	}

}
