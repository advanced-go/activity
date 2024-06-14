package assignment

import "time"

const (
	OpenStatus         = "open"
	ClosedStatus       = "closed"
	AssignedStatus     = "assigned"
	ReassignmentStatus = "reassignment"
)

// Case office looks for open assignments, and then does an assignment to a Service Agent and updating
// the assignment assignee id
// On reassignment - the assignee class and id are reset

// Entry - host
type Entry struct {
	EntryId   string    `json:"entry-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Origin + route for uniqueness
	Region  string `json:"region"`
	Zone    string `json:"zone"`
	SubZone string `json:"sub-zone"`
	Host    string `json:"host"`

	// Assignee class - these get reset, id = "", and class to new class
	AssigneeClass string `json:"assignee-class"`
	AssigneeId    string `json:"assignee-id"`
}

// When doing an assignment, the Agent id needs to be somewhere??

// EntryDetail - entry details
type EntryDetail struct {
	EntryId   string    `json:"entry-id"`
	DetailId  string    `json:"detail-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Details
	RouteName string `json:"status"`
	Details   string `json:"details"`
}

// EntryStatus - add an agentID?
type EntryStatus struct {
	EntryId   string    `json:"entry-id"`
	StatusId  string    `json:"status-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// New status and optional assignee id
	Status     string `json:"status"`
	AssigneeId string `json:"assignee-id"`
}

// EntryStatusUpdate - updates for reassignment and close
type EntryStatusUpdate struct {
	EntryId       string    `json:"entry-id"`
	UpdateId      string    `json:"update-id"`
	AgentId       string    `json:"agent-id"`
	CreatedTS     time.Time `json:"created-ts"`
	AssigneeClass string    `json:"assignee-class"`

	// Update data. Processed agent id needed ??
	// Error needed if updates are in an invalid order, such as a reassignment after a close
	NewStatus        string
	NewAssigneeClass string    `json:"new-assignee-class"`
	Error            string    `json:"error"`
	ProcessedTS      time.Time `json:"processed-ts"`
}

var entryData = []Entry{
	{EntryId: "1", AgentId: "test-agent", Region: "us-west", Zone: "oregon", Host: "www.host1.com", AssigneeClass: "case-officer:007", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: "2", AgentId: "test-agent", Region: "us-west", Zone: "oregon", Host: "www.host2.com", AssigneeClass: "case-officer:007", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
}

var entryDetailData = []EntryDetail{
	{EntryId: "1", DetailId: "1", AgentId: "agent-name:agent-class:instance-id", RouteName: "search", Details: "various information", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: "1", DetailId: "2", AgentId: "agent-name:agent-class:instance-id", RouteName: "host", Details: "other information", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
}

var entryStatusData = []EntryStatus{
	{EntryId: "1", StatusId: "1", AgentId: "agent-name:agent-class:instance-id", Status: "open", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: "1", StatusId: "2", AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
}
