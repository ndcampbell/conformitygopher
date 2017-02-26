package main

import (
	"log"

	"github.com/ndcampbell/conformitygopher/aws"
	"github.com/ndcampbell/conformitygopher/configs"
	"github.com/ndcampbell/conformitygopher/report"
)

func main() {
	log.Println("Starting ConformityGopher!")
	config := configs.LoadConfigs()
	aws.RunAll(config)
	if config.Email.Enabled {
		report.SendEmail(&config.Email)
	}
}
