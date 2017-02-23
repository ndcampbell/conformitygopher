package aws

import (
	"log"
	"sync"

	"github.com/ndcampbell/conformitygopher/configs"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

func RdsGather(sess *session.Session, rules *configs.RulesConfig, wg *sync.WaitGroup) {
	defer wg.Done()

	rdsclient := rds.New(sess)
	_, err := rdsclient.DescribeDBInstances(nil)
	if err != nil {
		log.Fatal("RDS Error", err)
	}
	log.Println("RDS Resources Gathered")
}
