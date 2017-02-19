package aws

import (
	"log"
	"sync"

	"github.com/ndcampbell/conformitygopher/configs"

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

func RunAll(config *configs.ConformityConfig) {
	var wg sync.WaitGroup
	for _, profile := range config.Profiles {
		log.Printf("Gathering for profile: %s", profile)
		sess := setupSession(profile)

		//Gather each resource in a goroutine, wait for all to complete before changing profiles
		wg.Add(1)
		go Ec2Gather(sess, &wg)
		wg.Add(1)
		go RdsGather(sess, &wg)

		wg.Wait()
	}
}
