package aws

import (
	"log"

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

func RunAll(config *configs.BaseConfig) *[]Resource {
	var badResources []Resource

	for _, profile := range config.Profiles {
		log.Printf("Gathering for profile: %s", profile)
		sess := setupSession(profile)

		for _, resource := range config.Resources {
			resourceFunc := ResourceMap[resource]
			tmpResource := resourceFunc(sess, &config.Rules)
			tmpResource.Account = profile
			badResources = append(badResources, tmpResource)
		}
	}
	return &badResources
}
