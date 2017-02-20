package aws

import (
	"log"
	"sync"

	"github.com/ndcampbell/conformitygopher/configs"
	"github.com/ndcampbell/conformitygopher/database"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func setupSession(profile string) *session.Session {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("us-east-1")},
		Profile: profile,
	})
	if err != nil {
		log.Fatal("AWS Session Error", err)
	}

	return sess
}

func RunAll() {
	var wg sync.WaitGroup

	database.DbSetup()
	for _, profile := range configs.Config.Profiles {
		log.Printf("Gathering for profile: %s", profile)
		sess := setupSession(profile)

		for _, resource := range configs.Config.Resources {
			resourceFunc := ResourceMap[resource]
			wg.Add(1)
			go resourceFunc(sess, &wg)
		}
		wg.Wait()
	}
}
