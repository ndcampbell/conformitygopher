package aws

import (
	"log"
	"sync"
	"time"

	"github.com/ndcampbell/conformitygopher/configs"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type InstanceData struct {
	InstanceId string
	Status     string
	LaunchTime time.Time
	BrokenRule string
}

func Ec2Gather(sess *session.Session, rules *configs.RulesConfig, wg *sync.WaitGroup) {
	defer wg.Done()

	ec2client := ec2.New(sess)
	resp, err := ec2client.DescribeInstances(nil)
	if err != nil {
		log.Fatal("EC2 Error", err)
	}
	log.Println("EC2 Resources Gathered")
	badInstances := iterateInstances(resp.Reservations, rules)
	log.Println(badInstances)
}

func iterateInstances(reservations []*ec2.Reservation, rules *configs.RulesConfig) []*InstanceData {
	var badInstances []*InstanceData
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

func checkRules(instance *ec2.Instance, rules *configs.RulesConfig) *InstanceData {
	var instanceData InstanceData
	tagRule := checkTags(instance.Tags, rules.RequiredTags)
	if tagRule == false {
		instanceData = buildInstanceData(instance, "Missing Required Tags")
		return &instanceData
	}
	return nil
}

func buildInstanceData(instance *ec2.Instance, brokenRule string) InstanceData {
	instanceData := InstanceData{
		InstanceId: *instance.InstanceId,
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
			log.Printf("%s Tag not found", requiredTag)
			return false
		}
	}
	return true
}
