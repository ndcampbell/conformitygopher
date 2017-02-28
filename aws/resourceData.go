package aws

import (
	"time"
)

type Resource struct {
	Type    string
	Account string
	Data    []*ResourceData
}

type ResourceData struct {
	Id         string
	Status     string
	LaunchTime time.Time
	BrokenRule string
}
