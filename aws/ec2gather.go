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

func Ec2Gather(sess *session.Session, wg *sync.WaitGroup) {
	defer wg.Done()

	ec2client := ec2.New(sess)
	resp, err := ec2client.DescribeInstances(nil)
	if err != nil {
		log.Fatal("EC2 Error", err)
	}
	log.Println("EC2 Resources Gathered")
	iterateInstances(resp.Reservations)

}

func iterateInstances(reservations []*ec2.Reservation) {
	var badInstances []*InstanceData
	for _, res := range reservations {
		for _, instance := range res.Instances {
			data := checkRules(instance)
			if data != nil {
				badInstances = append(badInstances, data)
			}
		}
	}
}

func checkRules(instance *ec2.Instance) *InstanceData {
	var instanceData InstanceData
	tagRule := checkTags(instance.Tags)
	if tagRule == false {
		instanceData = buildInstanceData(instance, "Missing Required Tags")
		log.Println(instanceData)
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

func checkTags(tags []*ec2.Tag) bool {
	for _, requiredTag := range configs.Config.Rules.RequiredTags {
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
