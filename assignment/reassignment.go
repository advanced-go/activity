package assignment

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

// getStatusChange - get status change by assignee for close
func getStatusChange(ctx context.Context, status, assigneeClass string, assigneeOrigin core.Origin) ([]EntryStatusChange, *core.Status) {
	//e, ok := index.LookupEntry(core.NewOrigin(values))
	//if !ok {
	//	return nil, core.StatusNotFound()
	//}
	defer safeChange.Lock()()
	for _, chg := range changeData {
		if chg.NewStatus != status {
			continue
		}
		if chg.AssigneeClass == assigneeClass && chg.AssigneeRegion == assigneeOrigin.Region && chg.AssigneeZone == assigneeOrigin.Zone && chg.AssigneeSubZone == assigneeOrigin.SubZone {
			return []EntryStatusChange{chg}, core.StatusOK()
		}
	}
	return nil, core.StatusNotFound()
}

// reassign - set the status of an assignment to reassignment, and update assignment receiver
func reassign(ctx context.Context, origin core.Origin, agentId, newAssigneeClass string, newAssigneeOrigin core.Origin) *core.Status {
	return core.StatusOK()
}

// processReassignment - process a reassignment
func processReassignment(ctx context.Context, change []EntryStatusChange) *core.Status {
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
