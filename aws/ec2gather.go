package aws

import (
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func Ec2Gather(sess *session.Session, wg *sync.WaitGroup) {
	defer wg.Done()

	ec2client := ec2.New(sess)
	_, err := ec2client.DescribeInstances(nil)
	if err != nil {
		log.Fatal("EC2 Error", err)
	}
	log.Println("EC2 Resources Gathered")
}
