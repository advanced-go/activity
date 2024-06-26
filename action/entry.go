package action

import (
	"errors"
	"fmt"
	"github.com/advanced-go/activity/common"
	"github.com/advanced-go/stdlib/core"
	"net/url"
	"time"
)

const (
	OpenStatus   = "open"
	ClosedStatus = "closed"
)

const (
	//accessLogSelect = "SELECT * FROM access_log {where} order by start_time limit 2"
	accessLogSelect = "SELECT region,customer_id,start_time,duration_str,traffic,rate_limit FROM access_log {where} order by start_time desc limit 2"
	accessLogInsert = "INSERT INTO access_log (" +
		"customer_id,start_time,duration_ms,duration_str,traffic," +
		"region,zone,sub_zone,service,instance_id,route_name," +
		"request_id,url,protocol,method,host,path,status_code,bytes_sent,status_flags," +
		"timeout,rate_limit,rate_burst) VALUES"
	deleteSql = "DELETE FROM access_log"

	EntryIdName    = "entry_id"
	AgentIdName    = "agent_id"
	CreatedTSName  = "created_ts"
	RegionName     = "region"
	ZoneName       = "zone"
	SubZoneName    = "sub_zone"
	HostName       = "host"
	RouteName      = "route"
	Action2Name    = "action"
	TimeoutName    = "timeout"
	RateLimitName  = "rate_limit"
	RateBurstName  = "rate_burst"
	PrimaryName    = "primary"
	SecondaryName  = "secondary"
	PercentageName = "percent"
	StatusName     = "status"
)

var (
	safeEntry = common.NewSafe()
	entryData = []Entry{
		{EntryId: 1, AgentId: "director-1", Region: "us-west-1", Zone: "usw1-az1", Host: "www.host1.com", Action: "test", Timeout: 0, RateLimit: 0, RateBurst: 0, Primary: "", Secondary: "", Percentage: 0, Status: OpenStatus, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 2, AgentId: "director-1", Region: "us-west-1", Zone: "usw1-az2", Host: "www.host2.com", Action: "test", Timeout: 0, RateLimit: 0, RateBurst: 0, Primary: "", Secondary: "", Percentage: 0, Status: OpenStatus, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 3, AgentId: "director-2", Region: "us-west-2", Zone: "usw2-az3", Host: "www.host1.com", Action: "test", Timeout: 0, RateLimit: 0, RateBurst: 0, Primary: "", Secondary: "", Percentage: 0, Status: OpenStatus, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 4, AgentId: "director-2", Region: "us-west-2", Zone: "usw2-az4", Host: "www.host2.com", Action: "test", Timeout: 0, RateLimit: 0, RateBurst: 0, Primary: "", Secondary: "", Percentage: 0, Status: OpenStatus, CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// Entry - host
type Entry struct {
	EntryId   int       `json:"entry-id"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	// Origin + route for uniqueness
	Region    string `json:"region"`
	Zone      string `json:"zone"`
	SubZone   string `json:"sub-zone"`
	Host      string `json:"host"`
	RouteName string `json:"route"`

	Action     string  `json:"action"`
	Timeout    int     `json:"timeout"`
	RateLimit  float64 `json:"rate-limit"`
	RateBurst  int     `json:"rate-burst"`
	Primary    string  `json:"primary"`
	Secondary  string  `json:"secondary"`
	Percentage int     `json:"percentage"`

	Status string
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
		case RouteName:
			e.RouteName = values[i].(string)

		case Action2Name:
			e.Action = values[i].(string)
		case TimeoutName:
			e.Timeout = values[i].(int)
		case RateLimitName:
			e.RateLimit = values[i].(float64)
		case RateBurstName:
			e.RateBurst = values[i].(int)
		case PrimaryName:
			e.Primary = values[i].(string)
		case SecondaryName:
			e.Secondary = values[i].(string)
		case PercentageName:
			e.Percentage = values[i].(int)
		case StatusName:
			e.Status = values[i].(string)
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
		e.Timeout,
		e.RateLimit,
		e.RateBurst,
		e.Primary,
		e.Secondary,
		e.Percentage,
		e.Status,
	}
}

func (Entry) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}

func validEntry(values url.Values, e Entry) bool {
	if values == nil {
		return false
	}
	filter := core.NewOrigin(values)
	target := core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}
	if !core.OriginMatch(target, filter) {
		return false
	}
	// Additional filtering
	return true
}
