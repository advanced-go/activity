package region

import "time"

const (
	StatusScheduled  = "scheduled"
	StatusProcessing = "processing"
	StatusCompleted  = "completed"
)

type Entry struct {
	EntryId   int       `json:"entry-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Assignment
	Region  string `json:"region"`
	Zone    string `json:"zone"`
	SubZone string `json:"sub-zone"`
}

type EntryStatus struct {
	EntryId   int       `json:"entry-id"`
	StatusId  int       `json:"status-id"`
	CreatedTS time.Time `json:"created-ts"`

	Status  string `json:"status"`
	Results string `json:"results"`
}

type EntryAssignment struct {
	EntryId      int       `json:"entry-id"`
	AssignmentId int       `json:"assignment-id"`
	CreatedTS    time.Time `json:"created-ts"`

	// Origin
	AssigneeTag string `json:"assignee-tag"`
}

type AssignmentStatus struct {
	AssignmentId int       `json:"assignment-id"`
	StatusId     int       `json:"status-id"`
	CreatedTS    time.Time `json:"created-ts"`

	Status string `json:"created-ts"`
}
