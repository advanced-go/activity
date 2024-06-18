package inference

import (
	"github.com/advanced-go/activity/common"
	"time"
)

// Entry - host
type Entry struct {
	EntryId   int       `json:"entry-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Origin + route name
	Region  string `json:"region"`
	Zone    string `json:"zone"`
	SubZone string `json:"sub-zone"`
	Host    string `json:"host"`

	// Details + action
	RouteName string `json:"route"`
	Details   string `json:"details"`
	Action    string `json:"action"`
}

var (
	safe      = common.NewSafe()
	entryData = []Entry{
		{EntryId: 1, AgentId: "agent", Region: "us-west", Zone: "oregon", Host: "www.host1.com", RouteName: "route", Details: "information", Action: "processed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 2, AgentId: "agent", Region: "us-west", Zone: "oregon", Host: "www.host2.com", RouteName: "host", Details: "text", Action: "processed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)
