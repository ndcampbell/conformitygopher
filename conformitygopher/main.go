package main

import (
	"log"

	"github.com/ndcampbell/conformitygopher/aws"
	"github.com/ndcampbell/conformitygopher/configs"
)

func main() {
	log.Println("Starting ConformityGopher!")
	config := configs.LoadConfigs()
	aws.RunAll(config)
}
