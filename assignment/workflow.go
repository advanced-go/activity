package assignment

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"time"
)

// insert - add an assignment and an open status
func insert(agentId string, origin core.Origin, assigneeTag string) *core.Status {
	// Enforce unique constraint
	_, ok := index.LookupEntry(origin)
	if ok {
		return core.NewStatusError(core.StatusInvalidContent, errors.New(fmt.Sprintf("error: assignment already exist for: %v", origin)))
	}
	defer safeEntry.Lock()()

	entry := Entry{
		EntryId:     lastEntry().EntryId + 1,
		AgentId:     agentId,
		CreatedTS:   time.Now().UTC(),
		Region:      origin.Region,
		Zone:        origin.Zone,
		SubZone:     origin.SubZone,
		Host:        origin.Host,
		AssigneeTag: assigneeTag,
		AssigneeId:  "",
		UpdatedTS:   time.Time{},
		Status:      "",
	}
	entryData = append(entryData, entry)
	index.AddEntry(entry)
	return addStatus(origin, OpenStatus, agentId, "")
}

// getOpen - find an open assignment for a given assignee tag
func getOpen(assigneeTag string) ([]Entry, *core.Status) {
	defer safeEntry.Lock()()
	defer safeStatus.Lock()()

	for _, e := range entryData {
		if e.AssigneeTag == assigneeTag {
			_, ok := lastStatusFilter(e.EntryId, OpenStatus)
			if ok {
				return []Entry{e}, core.StatusOK()
			}
		}
	}
	return nil, core.StatusNotFound()
}

func updateEntryStatus(origin core.Origin, status string) *core.Status {
	_, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeEntry.Lock()()

	for i, e := range entryData {
		if e.Region == origin.Region && e.Zone == origin.Zone && e.SubZone == origin.SubZone && e.Host == origin.Host {
			entryData[i].Status = status
		}
	}
	return core.StatusOK()
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
		DetailId:  lastDetail().DetailId + 1,
		AgentId:   agentId,
		CreatedTS: time.Now().UTC(),
		RouteName: routeName,
		Details:   details,
	}
	detailData = append(detailData, detail)
	return core.StatusOK()
}

// processClose - process a closed status update
func processClose(origin core.Origin, agentId string) *core.Status {
	_, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}

	status := addStatus(origin, ClosedStatus, agentId, "")
	if status.OK() {
		updateEntryStatus(origin, ClosedStatus)
	}
	return status
}

// processReassignment - process a reassignment
func processReassignment(origin core.Origin, agentId string, change EntryStatusChange) *core.Status {
	_, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}

	if change.NewStatus != ReassignedStatus {
		return core.StatusBadRequest()
	}

	status := addStatus(origin, ReassignedStatus, agentId, "")
	if status.OK() {
		updateEntryStatus(origin, ReassignedStatus)
	}
	return status
}

/*
func reassignEntry(o core.Origin, assigneeTag string) *core.Status {
	e, ok := index.LookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeEntry.Lock()()

	for i, entry := range entryData {
		if entry.EntryId == e.EntryId {
			entryData[i].UpdatedTS = time.Now().UTC()
			entryData[i].AssigneeTag = assigneeTag
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


*/
