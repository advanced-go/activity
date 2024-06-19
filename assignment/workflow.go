package assignment

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
	"time"
)

func getStatusChange(ctx context.Context, h http.Header, values url.Values) ([]EntryStatusChange, *core.Status) {
	e, ok := index.LookupEntry(core.NewOrigin(values))
	if !ok {
		return nil, core.StatusNotFound()
	}
	defer safeChange.Lock()()
	cls := ""
	s := values["assignee-class"]
	if len(s) > 0 {
		cls = s[0]
	}
	for _, change := range changeData {
		if change.EntryId == e.EntryId && change.AssigneeClass == cls {
			return []EntryStatusChange{change}, core.StatusOK()
		}
	}
	return nil, core.StatusNotFound()
}

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

func insertDetail(o core.Origin, detail EntryDetail) *core.Status {
	e, ok := index.LookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeDetail.Lock()()

	detail.EntryId = e.EntryId
	detail.DetailId = detailData[len(detailData)-1].DetailId + 1
	detail.CreatedTS = time.Now().UTC()
	detailData = append(detailData, detail)
	return core.StatusOK()
}

func insertStatus(o core.Origin, es EntryStatus) *core.Status {
	e, ok := index.LookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeStatus.Lock()()

	es.EntryId = e.EntryId
	es.StatusId = statusData[len(statusData)-1].StatusId + 1
	es.CreatedTS = time.Now().UTC()
	statusData = append(statusData, es)
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

func reassignEntry[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, assigneeClass string) *core.Status {
	e, ok := index.LookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeEntry.Lock()()

	for i, entry := range entryData {
		if entry.EntryId == e.EntryId {
			entryData[i].UpdatedTS = time.Now().UTC()
			entryData[i].AssigneeClass = assigneeClass
			//entryData[i].AssigneeId = ""
			return core.StatusOK()
		}
	}
	return core.StatusNotFound()
}

func assignEntry[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, assigneeId string) *core.Status {
	e, ok := index.LookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeEntry.Lock()()

	for i, entry := range entryData {
		if entry.EntryId == e.EntryId {
			entryData[i].UpdatedTS = time.Now().UTC()
			//entryData[i].AssigneeId = assigneeId
			return core.StatusOK()
		}
	}
	return core.StatusNotFound()
}
