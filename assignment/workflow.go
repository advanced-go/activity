package assignment

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

func GetEntryByStatus(cx context.Context, h http.Header, o core.Origin, status string) ([]Entry, *core.Status) {
	e, ok := lookupEntry(o)
	if !ok {
		return nil, core.StatusNotFound()
	}
	for i := len(statusData) - 1; i >= 0; i-- {
		if statusData[i].EntryId == e.EntryId && statusData[i].Status == status {
			return []Entry{e}, core.StatusOK()
		}
	}
	return nil, core.StatusNotFound()
}

func InsertEntry[E core.ErrorHandler](ctx context.Context, h http.Header, e Entry, assigneeId string) *core.Status {
	es := EntryStatus{
		EntryId:    e.EntryId,
		StatusId:   0,
		AgentId:    e.AgentId,
		CreatedTS:  time.Time{},
		Status:     OpenStatus,
		AssigneeId: assigneeId,
	}
	_, status := put[E, Entry](ctx, h, "", "", []Entry{e}, nil)
	if status.OK() {
		status = InsertStatus[E](ctx, h, e.Origin(), es)
	}
	return status
}

func ReassignEntry[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, assigneeClass string) *core.Status {
	e, ok := lookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	for i, entry := range entryData {
		if entry.EntryId == e.EntryId {
			entryData[i].UpdatedTS = time.Now().UTC()
			entryData[i].AssigneeClass = assigneeClass
			entryData[i].AssigneeId = ""
			return core.StatusOK()
		}
	}
	return core.StatusNotFound()
}

func AssignEntry[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, assigneeId string) *core.Status {
	e, ok := lookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	for i, entry := range entryData {
		if entry.EntryId == e.EntryId {
			entryData[i].UpdatedTS = time.Now().UTC()
			entryData[i].AssigneeId = assigneeId
			return core.StatusOK()
		}
	}
	return core.StatusNotFound()
}

func InsertDetail[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, detail EntryDetail) *core.Status {
	e, ok := lookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	detail.EntryId = e.EntryId
	detail.DetailId = detailData[len(detailData)-1].DetailId + 1
	_, status := put[E, EntryDetail](ctx, h, "", "", []EntryDetail{detail}, nil)
	return status
}

func InsertStatus[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, es EntryStatus) *core.Status {
	e, ok := lookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	es.EntryId = e.EntryId
	es.StatusId = statusData[len(statusData)-1].StatusId + 1
	_, status := put[E, EntryStatus](ctx, h, "", "", []EntryStatus{es}, nil)
	return status
}

func GetStatusChange[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, assigneeClass string) ([]EntryStatusChange, *core.Status) {
	e, ok := lookupEntry(o)
	if !ok {
		return nil, core.StatusNotFound()
	}
	for _, change := range changeData {
		if change.EntryId == e.EntryId && change.AssigneeClass == assigneeClass {
			return []EntryStatusChange{change}, core.StatusOK()
		}
	}
	return nil, core.StatusNotFound()
}

func InsertStatusChange[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, update EntryStatusChange) *core.Status {
	e, ok := lookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	update.EntryId = e.EntryId
	update.ChangeId = changeData[len(changeData)-1].ChangeId + 1
	_, status := put[E, EntryStatusChange](ctx, h, "", "", []EntryStatusChange{update}, nil)
	return status
}

func UpdateStatusChange[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, changeId int) *core.Status {
	e, ok := lookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	for i, u := range changeData {
		if u.EntryId == e.EntryId && u.ChangeId == changeId {
			changeData[i].UpdatedTS = time.Now().UTC()
			return core.StatusOK()
		}
	}
	//	_, status := post[E, EntryStatusChange](ctx, h, "", "", []EntryStatusChange{update}, nil)
	return core.StatusNotFound()
}
