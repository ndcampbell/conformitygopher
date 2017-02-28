package aws

import (
	"log"

	"github.com/ndcampbell/conformitygopher/configs"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func Ec2Gather(sess *session.Session, rules *configs.RulesConfig, c chan []*ResourceData) {

	ec2client := ec2.New(sess)
	resp, err := ec2client.DescribeInstances(nil)
	if err != nil {
		log.Fatal("EC2 Error", err)
	}
	log.Println("EC2 Resources Gathered")
	badInstances := iterateInstances(resp.Reservations, rules)
	c <- badInstances
}

func iterateInstances(reservations []*ec2.Reservation, rules *configs.RulesConfig) []*ResourceData {
	var badInstances []*ResourceData
	for _, res := range reservations {
		for _, instance := range res.Instances {
			data := checkRules(instance, rules)
			if data != nil {
				badInstances = append(badInstances, data)
			}
		}
	}
	return badInstances
}

func checkRules(instance *ec2.Instance, rules *configs.RulesConfig) *ResourceData {
	var instanceData ResourceData
	tagRule := checkTags(instance.Tags, rules.RequiredTags)
	if tagRule == false {
		instanceData = buildInstanceData(instance, "Missing Required Tags")
		return &instanceData
	}
	return nil
}

func buildInstanceData(instance *ec2.Instance, brokenRule string) ResourceData {
	instanceData := ResourceData{
		Id:         *instance.InstanceId,
		Status:     *instance.State.Name,
		LaunchTime: *instance.LaunchTime,
		BrokenRule: brokenRule,
	}
	return instanceData
}

func checkTags(tags []*ec2.Tag, requiredTags []string) bool {
	for _, requiredTag := range requiredTags {
		match := false
		for _, instanceTag := range tags {
			if requiredTag == *instanceTag.Key {
				match = true
				break
			}
		}
		if match == false {
			return false
		}
	}
	return true
}
