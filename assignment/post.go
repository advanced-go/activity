package assignment

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

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
			entryData[i].AssigneeId = ""
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
			entryData[i].AssigneeId = assigneeId
			return core.StatusOK()
		}
	}
	return core.StatusNotFound()
}

func updateStatusChange[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, changeId int) *core.Status {
	e, ok := index.LookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeChange.Lock()()

	for i, u := range changeData {
		if u.EntryId == e.EntryId && u.ChangeId == changeId {
			changeData[i].UpdatedTS = time.Now().UTC()
			return core.StatusOK()
		}
	}
	return core.StatusNotFound()
}
