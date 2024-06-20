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
		status = updateStatus(origin, ClosedStatus)
	}
	return status
}

// addStatus - add a status
func addStatus(origin core.Origin, status, agentId, assigneeId string) *core.Status {
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

func lastStatus(entryId int, status string) (EntryStatus, bool) {
	for i := len(statusData) - 1; i >= 0; i-- {
		if statusData[i].EntryId == entryId && statusData[i].Status == status {
			return statusData[i], true
		}
	}
	return EntryStatus{}, false
}

func addCloseStatusChange(origin core.Origin, agentId, assigneeTag string) *core.Status {
	e, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}
	// TODO: if current status is closing or closed then return an error status

	defer safeChange.Lock()()

	chg := EntryStatusChange{
		EntryId:        e.EntryId,
		ChangeId:       changeData[len(changeData)-1].ChangeId + 1,
		AgentId:        agentId,
		CreatedTS:      time.Now().UTC(),
		AssigneeTag:    assigneeTag,
		NewStatus:      ClosedStatus,
		NewAssigneeTag: "",
		Error:          "",
		UpdatedTS:      time.Time{},
	}
	changeData = append(changeData, chg)
	status := addStatus(origin, ClosingStatus, agentId, assigneeTag)
	if status.OK() {
		updateStatus(origin, ClosingStatus)
	}
	return core.StatusOK()
}

func addReassignmentStatusChange(origin core.Origin, agentId, assigneeTag, newAssigneeTag string) *core.Status {
	e, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}
	// TODO: if current status is reassigning or reassigned then return an error status

	defer safeChange.Lock()()

	chg := EntryStatusChange{
		EntryId:        e.EntryId,
		ChangeId:       changeData[len(changeData)-1].ChangeId + 1,
		AgentId:        agentId,
		CreatedTS:      time.Now().UTC(),
		AssigneeTag:    assigneeTag,
		NewStatus:      ReassignedStatus,
		NewAssigneeTag: newAssigneeTag,
		Error:          "",
		UpdatedTS:      time.Time{},
	}
	changeData = append(changeData, chg)
	status := addStatus(origin, ReassigningStatus, agentId, assigneeTag)
	if status.OK() {
		updateStatus(origin, ReassigningStatus)
	}
	return core.StatusOK()
}

// getStatusChange - get status change by assignee and class
func getStatusChange(status, assigneeTag string) (EntryStatusChange, *core.Status) {
	defer safeChange.Lock()()
	for _, chg := range changeData {
		if chg.NewStatus != status || chg.AssigneeTag != assigneeTag {
			continue
		}
		year, _, _ := chg.UpdatedTS.Date()
		if year == 1 {
			continue
		}
		return chg, core.StatusOK()
	}
	return EntryStatusChange{}, core.StatusNotFound()
}
