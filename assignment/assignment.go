package assignment

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"time"
)

// insert - add an assignment and an open status
func insert(ctx context.Context, agentId string, origin core.Origin, assigneeClass string, assigneeOrigin core.Origin) *core.Status {
	// Enforce unique constraint
	_, ok := index.LookupEntry(origin)
	if ok {
		return core.NewStatusError(core.StatusInvalidContent, errors.New(fmt.Sprintf("error: assignment already exist for: %v", origin)))
	}
	defer safeEntry.Lock()()

	entry := Entry{
		EntryId:         entryData[len(entryData)-1].EntryId + 10,
		AgentId:         agentId,
		CreatedTS:       time.Now().UTC(),
		Region:          origin.Region,
		Zone:            origin.Zone,
		SubZone:         origin.SubZone,
		Host:            origin.Host,
		AssigneeClass:   assigneeClass,
		AssigneeRegion:  assigneeOrigin.Region,
		AssigneeZone:    assigneeOrigin.Zone,
		AssigneeSubZone: assigneeOrigin.SubZone,
		AssigneeId:      "",
		UpdatedTS:       time.Time{},
		Status:          "",
	}
	entryData = append(entryData, entry)
	return addStatus(ctx, origin, OpenStatus, agentId, "")
}

// getOpen - find an open assignment for a given assignee class and origin
func getOpen(ctx context.Context, assigneeClass string, assigneeOrigin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}

// assign - set the status of an assignment to assigned
func assign(ctx context.Context, origin core.Origin, agentId, assigneeId string) *core.Status {
	return core.StatusOK()
}

// closeAssignment - add a closed status
func closeAssignment(ctx context.Context, origin core.Origin, agentId string) *core.Status {
	return addStatus(ctx, origin, ClosedStatus, agentId, "")
}

// addDetail - add assignment details
func addDetail(ctx context.Context, origin core.Origin, agentId, routeName, details string) *core.Status {
	e, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeDetail.Lock()()

	detail := EntryDetail{
		EntryId:   e.EntryId,
		DetailId:  detailData[len(detailData)-1].DetailId + 1,
		AgentId:   agentId,
		CreatedTS: time.Now().UTC(),
		RouteName: routeName,
		Details:   details,
	}
	detailData = append(detailData, detail)
	return core.StatusOK()
}

// addStatus - update the status of an assignment to closed
func addStatus(ctx context.Context, origin core.Origin, status, agentId, assigneeId string) *core.Status {
	e, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeStatus.Lock()()

	es := EntryStatus{
		EntryId:    e.EntryId,
		StatusId:   statusData[len(statusData)-1].StatusId + 1,
		AgentId:    agentId,
		CreatedTS:  time.Now().UTC(),
		Status:     status,
		AssigneeId: assigneeId,
	}
	statusData = append(statusData, es)
	return core.StatusOK()
}
