package common

import (
	"fmt"
	"github.com/advanced-go/common/uri"
	"time"
)

// EntryStatus - add an agentID?
type EntryStatus struct {
	EntryId   int       `json:"entry-id"`
	StatusId  int       `json:"status-id"`
	AgentId   string    `json:"agent-id"` // Creation agent id
	CreatedTS time.Time `json:"created-ts"`

	// New status and optional assignee id
	Status     string `json:"status"`
	AssigneeId string `json:"assignee-id"` // Used to set assigned agent id when status is assigned
}

func (e EntryStatus) GetEntryId() int {
	return e.EntryId
}

func (e EntryStatus) GetStatus() string {
	return e.Status
}

var statusData = []EntryStatus{
	{EntryId: 1, StatusId: 1, AgentId: "agent-name:agent-class:instance-id", Status: "open", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: 1, StatusId: 2, AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: 3, StatusId: 3, AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: 4, StatusId: 4, AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
}

func ExampleValidStatus() {
	q := "entry-id=5"
	status := ValidStatus(uri.BuildValues(q), statusData[0])
	fmt.Printf("test: ValidStatus(\"%v\") -> [status:%v]\n", q, status)

	q = "entry-id=1"
	status = ValidStatus(uri.BuildValues(q), statusData[0])
	fmt.Printf("test: ValidStatus(\"%v\") -> [status:%v]\n", q, status)

	q = "entry-id=1&status=closed"
	status = ValidStatus(uri.BuildValues(q), statusData[0])
	fmt.Printf("test: ValidStatus(\"%v\") -> [status:%v]\n", q, status)

	q = "entry-id=1&status=open"
	status = ValidStatus(uri.BuildValues(q), statusData[0])
	fmt.Printf("test: ValidStatus(\"%v\") -> [status:%v]\n", q, status)

	//Output:
	//test: ValidStatus("entry-id=5") -> [status:false]
	//test: ValidStatus("entry-id=1") -> [status:true]
	//test: ValidStatus("entry-id=1&status=closed") -> [status:false]
	//test: ValidStatus("entry-id=1&status=open") -> [status:true]

}
