package aws

import (
	"github.com/ndcampbell/conformitygopher/configs"

	"github.com/aws/aws-sdk-go/aws/session"
)

/*
map of AWS gather functions. Maps to a string that would be found in
JSON configs. As long as all gathers have the same function interface, this is
easy to maintain and loop over
*/

var ResourceMap = map[string]func(*session.Session, *configs.RulesConfig, chan Resource){
	"ec2": Ec2Gather,
	"rds": RdsGather,
	"elb": ElbGather,
}
