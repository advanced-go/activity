package inference

import (
	"errors"
	"fmt"
	"github.com/advanced-go/activity/common"
	"github.com/advanced-go/stdlib/core"
	"time"
)

const (
	//accessLogSelect = "SELECT * FROM access_log {where} order by start_time limit 2"
	accessLogSelect = "SELECT region,customer_id,start_time,duration_str,traffic,rate_limit FROM access_log {where} order by start_time desc limit 2"
	accessLogInsert = "INSERT INTO access_log (" +
		"customer_id,start_time,duration_ms,duration_str,traffic," +
		"region,zone,sub_zone,service,instance_id,route_name," +
		"request_id,url,protocol,method,host,path,status_code,bytes_sent,status_flags," +
		"timeout,rate_limit,rate_burst) VALUES"

	EntryIdName   = "entry_id"
	AgentIdName   = "agent_id"
	CreatedTSName = "created_ts"
	RegionName    = "region"
	ZoneName      = "zone"
	SubZoneName   = "sub_zone"
	HostName      = "host"
	RouteName     = "route"
	DetailsName   = "details"
	ActionName    = "action"
)

var (
	safe      = common.NewSafe()
	entryData = []Entry{
		{EntryId: 1, AgentId: "agent", Region: "us-west", Zone: "oregon", Host: "www.host1.com", RouteName: "route", Details: "information", Action: "processed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 2, AgentId: "agent", Region: "us-west", Zone: "oregon", Host: "www.host2.com", RouteName: "host", Details: "text", Action: "processed", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// Entry - host
type Entry struct {
	EntryId   int       `json:"entry-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Origin + route name
	Region    string `json:"region"`
	Zone      string `json:"zone"`
	SubZone   string `json:"sub-zone"`
	Host      string `json:"host"`
	RouteName string `json:"route"`

	// Details + action
	Details string `json:"details"`
	Action  string `json:"action"`
}

func (e Entry) Origin() core.Origin {
	return core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}
}

func (Entry) Scan(columnNames []string, values []any) (e Entry, err error) {
	for i, name := range columnNames {
		switch name {
		case EntryIdName:
			e.EntryId = values[i].(int)
		case AgentIdName:
			e.AgentId = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)

		case RegionName:
			e.Region = values[i].(string)
		case ZoneName:
			e.Zone = values[i].(string)
		case SubZoneName:
			e.SubZone = values[i].(string)
		case HostName:
			e.Host = values[i].(string)

		case ActionName:
			e.Action = values[i].(string)
		case DetailsName:
			e.Details = values[i].(string)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e Entry) Values() []any {
	return []any{
		e.EntryId,
		e.AgentId,
		e.CreatedTS,

		e.Region,
		e.Zone,
		e.SubZone,
		e.Host,
		e.RouteName,

		e.Action,
		e.Details,
	}
}

func (Entry) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
