package aws

import (
	"log"

	"github.com/ndcampbell/conformitygopher/configs"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
)

func ElbGather(sess *session.Session, rules *configs.RulesConfig) Resource {

	elbclient := elb.New(sess)
	resp, err := elbclient.DescribeLoadBalancers(nil)
	if err != nil {
		log.Fatal("ELB Error", err)
	}
	log.Println("ELB Resources Gathered")
	iterateElbs(resp.LoadBalancerDescriptions, rules)
	return Resource{}
}

func iterateElbs(elbs []*elb.LoadBalancerDescription, rules *configs.RulesConfig) {
	for _, lb := range elbs {
		log.Println(lb)
	}
}
