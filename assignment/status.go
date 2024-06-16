package assignment

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

const (

	//accessLogSelect = "SELECT * FROM access_log {where} order by start_time limit 2"
	statusSelect = "SELECT region,customer_id,start_time,duration_str,traffic,rate_limit FROM access_log {where} order by start_time desc limit 2"

	statusInsert = "INSERT INTO access_log (" +
		"customer_id,start_time,duration_ms,duration_str,traffic," +
		"region,zone,sub_zone,service,instance_id,route_name," +
		"request_id,url,protocol,method,host,path,status_code,bytes_sent,status_flags," +
		"timeout,rate_limit,rate_burst,retry,retry_rate_limit,retry_rate_burst,failover) VALUES"

	//deleteSql = "DELETE FROM access_log"

	StatusIdName         = "status_id"
	UpdateIdName         = "update_id"
	StatusName           = "status"
	NewStatusName        = "new_status"
	NewAssigneeClassName = "new_assignee_class"
	ErrorName            = "error"
	ProcessedTSName      = "processed_ts"
)

var statusData = []EntryStatus{
	{EntryId: 1, StatusId: 1, AgentId: "agent-name:agent-class:instance-id", Status: "open", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: 1, StatusId: 2, AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: 3, StatusId: 3, AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: 4, StatusId: 4, AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
}

// EntryStatus - add an agentID?
type EntryStatus struct {
	EntryId   int       `json:"entry-id"`
	StatusId  int       `json:"status-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// New status and optional assignee id
	Status     string `json:"status"`
	AssigneeId string `json:"assignee-id"`
}

func (EntryStatus) Scan(columnNames []string, values []any) (e EntryStatus, err error) {
	for i, name := range columnNames {
		switch name {
		case EntryIdName:
			e.EntryId = values[i].(int)
		case StatusIdName:
			e.StatusId = values[i].(int)
		case AgentIdName:
			e.AgentId = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)
		case StatusName:
			e.Status = values[i].(string)
		case AssigneeIdName:
			e.AssigneeId = values[i].(string)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (a EntryStatus) Values() []any {
	return []any{
		a.EntryId,
		a.StatusId,
		a.AgentId,
		a.CreatedTS,
		a.Status,
		a.AssigneeId,
	}
}

func (EntryStatus) Rows(entries []EntryStatus) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}

var updateData = []EntryStatusChange{
	{EntryId: 1, UpdateId: 1, AgentId: "agent-name:agent-class:instance-id", AssigneeClass: "class", NewStatus: "closed", NewAssigneeClass: "new", Error: "test error", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: 1, UpdateId: 2, AgentId: "agent-name:agent-class:instance-id", AssigneeClass: "class2", NewStatus: "closed", NewAssigneeClass: "new", Error: "test2 error", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
}

// EntryStatusChange - updates for reassignment and close
type EntryStatusChange struct {
	EntryId       int       `json:"entry-id"`
	UpdateId      int       `json:"update-id"`
	AgentId       string    `json:"agent-id"`
	CreatedTS     time.Time `json:"created-ts"`
	AssigneeClass string    `json:"assignee-class"`

	// Update data. Processed agent id needed ??
	// Error needed if updates are in an invalid order, such as a reassignment after a close
	NewStatus        string    `json:"new-status"`
	NewAssigneeClass string    `json:"new-assignee-class"`
	Error            string    `json:"error"`
	UpdatedTS        time.Time `json:"updated-ts"`
}

func (EntryStatusChange) Scan(columnNames []string, values []any) (e EntryStatusChange, err error) {
	for i, name := range columnNames {
		switch name {
		case EntryIdName:
			e.EntryId = values[i].(int)
		case UpdateIdName:
			e.UpdateId = values[i].(int)
		case AgentIdName:
			e.AgentId = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)
		case AssigneeClassName:
			e.AssigneeClass = values[i].(string)
		case NewStatusName:
			e.NewStatus = values[i].(string)
		case NewAssigneeClassName:
			e.NewAssigneeClass = values[i].(string)
		case ErrorName:
			e.Error = values[i].(string)
		case ProcessedTSName:
			e.UpdatedTS = values[i].(time.Time)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (a EntryStatusChange) Values() []any {
	return []any{
		a.EntryId,
		a.UpdateId,
		a.AgentId,
		a.CreatedTS,
		a.AssigneeClass,
		a.NewStatus,
		a.NewAssigneeClass,
		a.Error,
		a.UpdatedTS,
	}
}

func (EntryStatusChange) Rows(entries []EntryStatusChange) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}

func validStatus(values url.Values, e EntryStatus) bool {
	if values == nil || values.Get("entry-id") != string(e.EntryId) {
		return false
	}

	// Additional filtering
	return true
}

func validStatusUpdate(values url.Values, e EntryStatusChange) bool {
	if values == nil || values.Get("entry-id") != string(e.EntryId) {
		return false
	}

	// Additional filtering
	return true
}
