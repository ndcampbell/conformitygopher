package database

import (
	"log"

	"github.com/ndcampbell/conformitygopher/configs"
)

func DbSetup(config *configs.ConformityConfig) {
    if config.Db.Type == "in-memory" {
        BoltSetup(config.Db.Location)
        log.Println("Bolt DB Setup")
    }

}
