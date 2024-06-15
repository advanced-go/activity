package assignment

import (
	"errors"
	"fmt"
	"time"
)

const (
	detailSelect = "SELECT * FROM access_log {where} order by start_time limit 2"
	//accessLogSelect = "SELECT region,customer_id,start_time,duration_str,traffic,rate_limit FROM access_log {where} order by start_time desc limit 2"

	detailInsert = "INSERT INTO access_log (" +
		"customer_id,start_time,duration_ms,duration_str,traffic," +
		"region,zone,sub_zone,service,instance_id,route_name," +
		"request_id,url,protocol,method,host,path,status_code,bytes_sent,status_flags," +
		"timeout,rate_limit,rate_burst,retry,retry_rate_limit,retry_rate_burst,failover) VALUES"

	DetailIdName  = "detail_id"
	RouteNameName = "route"
	DetailsName   = "details"
)

var entryDetailData = []EntryDetail{
	{EntryId: "1", DetailId: "1", AgentId: "agent-name:agent-class:instance-id", RouteName: "search", Details: "various information", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	{EntryId: "1", DetailId: "2", AgentId: "agent-name:agent-class:instance-id", RouteName: "host", Details: "other information", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
}

// EntryDetail - entry details
type EntryDetail struct {
	EntryId   string    `json:"entry-id"`
	DetailId  string    `json:"detail-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Details
	RouteName string `json:"route"`
	Details   string `json:"details"`
}

func (EntryDetail) Scan(columnNames []string, values []any) (e EntryDetail, err error) {
	for i, name := range columnNames {
		switch name {
		case EntryIdName:
			e.EntryId = values[i].(string)
		case DetailIdName:
			e.DetailId = values[i].(string)
		case AgentIdName:
			e.AgentId = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)
		case RouteNameName:
			e.RouteName = values[i].(string)
		case DetailsName:
			e.Details = values[i].(string)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e EntryDetail) Values() []any {
	return []any{
		e.EntryId,
		e.DetailId,
		e.AgentId,
		e.CreatedTS,
		e.RouteName,
		e.Details,
	}
}

func (EntryDetail) CreateInsertValues(entries []EntryDetail) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
