package assignment

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"time"
)

// insert - add an assignment and an open status
func insert(agentId string, origin core.Origin, assigneeClass string, assigneeOrigin core.Origin) *core.Status {
	// Enforce unique constraint
	_, ok := index.LookupEntry(origin)
	if ok {
		return core.NewStatusError(core.StatusInvalidContent, errors.New(fmt.Sprintf("error: assignment already exist for: %v", origin)))
	}
	defer safeEntry.Lock()()

	entry := Entry{
		EntryId:         entryData[len(entryData)-1].EntryId + 1,
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
	index.AddEntry(entry)
	return addStatus(origin, OpenStatus, agentId, "")
}

// getOpen - find an open assignment for a given assignee class and origin
func getOpen(assigneeClass string, assigneeOrigin core.Origin) ([]Entry, *core.Status) {
	defer safeEntry.Lock()()
	defer safeStatus.Lock()()

	for _, e := range entryData {
		if e.AssigneeClass != assigneeClass {
			continue
		}
		if e.AssigneeRegion == assigneeOrigin.Region && e.AssigneeZone == assigneeOrigin.Zone && e.AssigneeSubZone == assigneeOrigin.SubZone {
			_, ok := lastStatus(e.EntryId, OpenStatus)
			if ok {
				return []Entry{e}, core.StatusOK()
			}
		}
	}
	return nil, core.StatusNotFound()
}

// addDetail - add assignment details
func addDetail(origin core.Origin, agentId, routeName, details string) *core.Status {
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

// processReassignment - process a reassignment
func processReassignment(ctx context.Context, change []EntryStatusChange) *core.Status {
	return core.StatusOK()
}

func reassignEntry(o core.Origin, assigneeClass string) *core.Status {
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

func assignEntry(o core.Origin, assigneeId string) *core.Status {
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
