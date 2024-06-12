package inference

import "time"

// Entry - host
type Entry struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	CreatedTS time.Time `json:"created-ts"`
	Agent     string    `json:"agent"`
	Action    string    `json:"action"`
}

var storage = []Entry{
	{Region: "us-west", Zone: "oregon", SubZone: "dc1", Host: "www.host1.com", Agent: "test-agent", Action: "processed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{Region: "us-west", Zone: "oregon", SubZone: "dc2", Host: "www.host2.com", Agent: "test-agent", Action: "processed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
}
