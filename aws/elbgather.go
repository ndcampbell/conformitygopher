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
	resource := Resource{Type: "ELB"}
	log.Println("ELB Resources Gathered")
	resource.Data = iterateElbs(resp.LoadBalancerDescriptions, rules)
	return resource
}

func iterateElbs(elbs []*elb.LoadBalancerDescription, rules *configs.RulesConfig) []*ResourceData {
	var badElbs []*ResourceData
	for _, lb := range elbs {
		checkLbRules(lb, rules)
	}
	return badElbs
}

func checkLbRules(lb *elb.LoadBalancerDescription, rules *configs.RulesConfig) *ResourceData {
	var lbData ResourceData
	if rules.EmptyElb == true {
		if len(lb.Instances) == 0 {
			lbData = buildLbData(lb, "Empty Load Balancer")
			return &lbData
		}
	}
	return nil
}

func buildLbData(lb *elb.LoadBalancerDescription, brokenRule string) ResourceData {
	lbData := ResourceData{
		Id:         *lb.LoadBalancerName,
		Status:     " ",
		LaunchTime: *lb.CreatedTime,
		BrokenRule: brokenRule,
	}
	return lbData
}
