package aws

import (
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type InstanceData struct {
	InstanceId string
	Status     string
	LaunchTime time.Time
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
	var badInstances []InstanceData
	for _, res := range reservations {
		for _, instance := range res.Instances {
			data := checkRules(instance)
			badInstances = append(badInstances, data)
		}
	}

	log.Println(badInstances)
}

func checkRules(instance *ec2.Instance) InstanceData {
	instanceData := InstanceData{
		InstanceId: *instance.InstanceId,
		Status:     *instance.State.Name,
		LaunchTime: *instance.LaunchTime,
	}
	return instanceData

}
