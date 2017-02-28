package aws

import (
	"log"

	"github.com/ndcampbell/conformitygopher/configs"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

func RdsGather(sess *session.Session, rules *configs.RulesConfig, c chan []*ResourceData) {

	rdsclient := rds.New(sess)

	log.Println(c) //placeholder

	_, err := rdsclient.DescribeDBInstances(nil)
	if err != nil {
		log.Fatal("RDS Error", err)
	}
	log.Println("RDS Resources Gathered")
}
