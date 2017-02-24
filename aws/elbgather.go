package aws

import (
	"log"
	"sync"
	"time"

	"github.com/ndcampbell/conformitygopher/configs"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
)

type ElbData struct {
	LoadBalancerName string
	LaunchTime       time.Time
	BrokenRule       string
}

func ElbGather(sess *session.Session, rules *configs.RulesConfig, wg *sync.WaitGroup) {
	defer wg.Done()

	elbclient := elb.New(sess)
	resp, err := elbclient.DescribeLoadBalancers(nil)
	if err != nil {
		log.Fatal("ELB Error", err)
	}
	log.Println("ELB Resources Gathered")
	iterateElbs(resp.LoadBalancerDescriptions, rules)
}

func iterateElbs(elbs []*elb.LoadBalancerDescription, rules *configs.RulesConfig) {
	for _, lb := range elbs {
		log.Println(lb)
	}
}
