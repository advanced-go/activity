package assignment

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

// assign - add an assigned status
func assign(origin core.Origin, agentId, assigneeId string) *core.Status {
	_, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}

	status := addStatus(origin, AssignedStatus, agentId, assigneeId)
	if status.OK() {
		status = updateEntryStatus(origin, AssignedStatus)
	}
	return status
}

func addClosingStatusChange(origin core.Origin, agentId, assigneeTag string) *core.Status {
	e, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}
	// TODO: if current status is closing or closed then return an error status

	defer safeChange.Lock()()

	chg := EntryStatusChange{
		EntryId:        e.EntryId,
		ChangeId:       lastChange().ChangeId + 1,
		AgentId:        agentId,
		CreatedTS:      time.Now().UTC(),
		AssigneeTag:    assigneeTag,
		NewStatus:      ClosingStatus,
		NewAssigneeTag: "",
		Error:          "",
		UpdatedTS:      time.Time{},
	}
	changeData = append(changeData, chg)
	status := addStatus(origin, ClosingStatus, agentId, assigneeTag)
	if status.OK() {
		status = updateEntryStatus(origin, ClosingStatus)
	}
	return status
}

func addReassigningStatusChange(origin core.Origin, agentId, assigneeTag, newAssigneeTag string) *core.Status {
	e, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}
	// TODO: if current status is reassigning or reassigned then return an error status

	defer safeChange.Lock()()

	chg := EntryStatusChange{
		EntryId:        e.EntryId,
		ChangeId:       lastChange().ChangeId + 1,
		AgentId:        agentId,
		CreatedTS:      time.Now().UTC(),
		AssigneeTag:    assigneeTag,
		NewStatus:      ReassigningStatus,
		NewAssigneeTag: newAssigneeTag,
		Error:          "",
		UpdatedTS:      time.Time{},
	}
	changeData = append(changeData, chg)
	status := addStatus(origin, ReassigningStatus, agentId, assigneeTag)
	if status.OK() {
		status = updateEntryStatus(origin, ReassigningStatus)
	}
	return status
}

// getStatusChange - get status change by assignee and class
func getStatusChange(status, assigneeTag string) (EntryStatusChange, *core.Status) {
	defer safeChange.Lock()()
	for _, chg := range changeData {
		if chg.NewStatus != status || chg.AssigneeTag != assigneeTag {
			continue
		}
		year, _, _ := chg.UpdatedTS.Date()
		if year != 1 {
			continue
		}
		return chg, core.StatusOK()
	}
	return EntryStatusChange{}, core.StatusNotFound()
}

// addStatus - add a status
func addStatus(origin core.Origin, status, agentId, assigneeId string) *core.Status {
	e, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}
	if status == "" {
		return core.StatusBadRequest()
	}
	defer safeStatus.Lock()()

	es := EntryStatus{
		EntryId:    e.EntryId,
		StatusId:   lastStatus().StatusId + 1,
		AgentId:    agentId,
		CreatedTS:  time.Now().UTC(),
		Status:     status,
		AssigneeId: assigneeId,
	}
	statusData = append(statusData, es)
	return core.StatusOK()
}

func lastStatusFilter(entryId int, status string) (EntryStatus, bool) {
	if entryId < 0 || status == "" {
		return EntryStatus{}, false
	}
	for i := len(statusData) - 1; i >= 0; i-- {
		if statusData[i].EntryId == entryId && statusData[i].Status == status {
			return statusData[i], true
		}
	}
	return EntryStatus{}, false
}
