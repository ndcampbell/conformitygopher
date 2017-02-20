package main

import (
	"log"

	"github.com/ndcampbell/conformitygopher/aws"
	"github.com/ndcampbell/conformitygopher/configs"
)

func main() {
	log.Println("Starting ConformityGopher!")
	configs.LoadConfigs()
	aws.RunAll()
}
