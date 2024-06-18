package common

import (
	"github.com/advanced-go/stdlib/core"
	"net/url"
	"time"
)

// Entry - host
type Entry struct {
	EntryId   int       `json:"entry-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Origin + route for uniqueness
	Region  string `json:"region"`
	Zone    string `json:"zone"`
	SubZone string `json:"sub-zone"`
	Host    string `json:"host"`
}

func (e Entry) Origin() core.Origin {
	return core.Origin{Region: e.Region, Zone: e.Zone, Host: e.Host}
}

var entryData = []Entry{
	{EntryId: 1, AgentId: "director-1", Region: "us-west-1", Zone: "usw1-az1", Host: "www.host1.com", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: 2, AgentId: "director-1", Region: "us-west-1", Zone: "usw1-az2", Host: "www.host2.com", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: 3, AgentId: "director-2", Region: "us-west-2", Zone: "usw2-az3", Host: "www.host1.com", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: 4, AgentId: "director-2", Region: "us-west-2", Zone: "usw2-az4", Host: "www.host2.com", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
}

var entryList = NewOriginIndex[Entry](entryData)

func ValidEntry(values url.Values, e Entry) bool {
	if values == nil {
		return false
	}
	filter := core.NewOrigin(values)
	if !core.OriginMatch(e.Origin(), filter) {
		return false
	}
	// Additional filtering
	return true
}
