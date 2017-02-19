package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func Ec2Gather(sess *session.Session) {
	ec2client := ec2.New(sess)
	_, err := ec2client.DescribeInstances(nil)
	if err != nil {
		log.Fatal("EC2 Error", err)
	}
}
