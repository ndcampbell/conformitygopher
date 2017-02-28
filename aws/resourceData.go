package aws

import (
	"time"
)

type ResourceData struct {
	Id         string
	Status     string
	LaunchTime time.Time
	BrokenRule string
}
