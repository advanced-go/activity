package action

import (
	"errors"
	"fmt"
	"github.com/advanced-go/activity/common"
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

	StatusIdName = "status_id"
	StatusName   = "status"
)

var (
	statusList = common.NewSafeSlice[EntryStatus](statusData)

	statusData = []EntryStatus{
		{EntryId: 1, StatusId: 1, AgentId: "agent-name:agent-class:instance-id", Status: "open", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 1, StatusId: 2, AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 3, StatusId: 3, AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 4, StatusId: 4, AgentId: "agent-name:agent-class:instance-id", Status: "closed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// EntryStatus - status
type EntryStatus struct {
	EntryId   int       `json:"entry-id"`
	StatusId  int       `json:"status-id"`
	AgentId   string    `json:"agent-id"` // Creation agent id
	CreatedTS time.Time `json:"created-ts"`

	Status string `json:"status"`
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

	s := values.Get("status")
	if s != "" && e.Status != s {
		return false
	}
	return true
}
