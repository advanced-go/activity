package audit

import "time"

// Entry - host
type Entry struct {
	EntryId   int       `json:"entry-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Origin
	Region  string `json:"region"`
	Zone    string `json:"zone"`
	SubZone string `json:"sub-zone"`
	Host    string `json:"host"`

	// Region + Zone + Class
	AssigneeTag string `json:"assignee-tag"` // Assigned to an agent class and origin
	AssigneeId  string `json:"assignee-id"`  // Set when an agent pulls this entry

	UpdatedTS time.Time `json:"updated-ts"`
	Status    string    `json:"status"`
}
