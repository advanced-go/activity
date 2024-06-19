package assignment

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

func NewOrigin(values map[string][]string) core.Origin {
	region := ""
	zone := ""
	subZone := ""
	host := ""

	s := values[core.RegionKey]
	if len(s) > 0 {
		region = s[0]
	}
	s = values[core.ZoneKey]
	if len(s) > 0 {
		zone = s[0]
	}
	s = values[core.SubZoneKey]
	if len(s) > 0 {
		subZone = s[0]
	}
	s = values[core.HostKey]
	if len(s) > 0 {
		host = s[0]
	}
	return core.Origin{Region: region, Zone: zone, SubZone: subZone, Host: host}
}

func insertEntry(e Entry, assigneeId string) *core.Status {
	defer safeEntry.Lock()()

	es := EntryStatus{EntryId: e.EntryId, StatusId: 0, AgentId: e.AgentId, CreatedTS: time.Time{}, Status: OpenStatus, AssigneeId: assigneeId}
	e.CreatedTS = time.Now().UTC()
	e.EntryId = entryData[len(entryData)-1].EntryId + 1
	insertStatus(e.Origin(), es)
	return core.StatusOK()
}

func insertStatusChange(o core.Origin, change EntryStatusChange) *core.Status {
	e, ok := index.LookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeChange.Lock()()

	change.EntryId = e.EntryId
	change.ChangeId = changeData[len(changeData)-1].ChangeId + 1
	change.CreatedTS = time.Now().UTC()
	changeData = append(changeData, change)
	return core.StatusOK()
}
