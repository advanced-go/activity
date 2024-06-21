package assignment

import (
	"errors"
	"fmt"
	"github.com/advanced-go/activity/common"
	"net/url"
	"time"
)

const (
	OpenStatus        = "open"
	AssignedStatus    = "assigned"
	ReassigningStatus = "reassigning"
	ClosingStatus     = "closing"
	ClosedStatus      = "closed"
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

	StatusIdName = "status_id"
	ChangeIdName = "change_id"
	StatusName   = "status"
)

var (
	safeStatus = common.NewSafe()
	statusData = []EntryStatus{
		{EntryId: 1, StatusId: 1, AgentId: "agent-name:agent-class:instance-id", Status: "open", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 1, StatusId: 2, AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 3, StatusId: 3, AgentId: "agent-name:agent-class:instance-id", Status: "closing", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 4, StatusId: 4, AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

func lastStatus() EntryStatus {
	return statusData[len(statusData)-1]
}

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

func validStatus(values url.Values, e EntryStatus) bool {
	if values == nil || values.Get("entry-id") != fmt.Sprintf("%v", e.EntryId) {
		return false
	}

	// Additional filtering
	return true
}
