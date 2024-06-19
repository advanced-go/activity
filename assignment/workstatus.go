package assignment

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

// assign - add an assigned status
func assign(origin core.Origin, agentId, assigneeId string) *core.Status {
	return addStatus(origin, AssignedStatus, agentId, assigneeId)
}

// closeAssignment - add a closed status
func closeAssignment(origin core.Origin, agentId string) *core.Status {
	return addStatus(origin, ClosedStatus, agentId, "")
}

// reassign - set the status of an assignment to reassignment, and update assignment receiver
func reassign(origin core.Origin, agentId, newAssigneeClass string, newAssigneeOrigin core.Origin) *core.Status {
	return core.StatusOK()
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

func addCloseStatusChange(origin core.Origin, agentId, assigneeClass string, assigneeOrigin core.Origin) *core.Status {
	e, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeChange.Lock()()

	chg := EntryStatusChange{
		EntryId:            e.EntryId,
		ChangeId:           changeData[len(changeData)-1].ChangeId + 1,
		AgentId:            agentId,
		CreatedTS:          time.Now().UTC(),
		AssigneeClass:      assigneeClass,
		AssigneeRegion:     assigneeOrigin.Region,
		AssigneeZone:       assigneeOrigin.Zone,
		AssigneeSubZone:    assigneeOrigin.SubZone,
		NewStatus:          ClosedStatus,
		NewAssigneeClass:   "",
		NewAssigneeRegion:  "",
		NewAssigneeZone:    "",
		NewAssigneeSubZone: "",
		Error:              "",
		UpdatedTS:          time.Time{},
	}
	changeData = append(changeData, chg)
	return core.StatusOK()
}

func addReassignmentStatusChange(origin core.Origin, agentId, assigneeClass string, assigneeOrigin core.Origin, newAssigneeClass string, newAssigneeOrigin core.Origin) *core.Status {
	e, ok := index.LookupEntry(origin)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeChange.Lock()()

	chg := EntryStatusChange{
		EntryId:            e.EntryId,
		ChangeId:           changeData[len(changeData)-1].ChangeId + 1,
		AgentId:            agentId,
		CreatedTS:          time.Now().UTC(),
		AssigneeClass:      assigneeClass,
		AssigneeRegion:     assigneeOrigin.Region,
		AssigneeZone:       assigneeOrigin.Zone,
		AssigneeSubZone:    assigneeOrigin.SubZone,
		NewStatus:          ReassignmentStatus,
		NewAssigneeClass:   newAssigneeClass,
		NewAssigneeRegion:  newAssigneeOrigin.Region,
		NewAssigneeZone:    newAssigneeOrigin.Zone,
		NewAssigneeSubZone: newAssigneeOrigin.SubZone,
		Error:              "",
		UpdatedTS:          time.Time{},
	}
	changeData = append(changeData, chg)
	return core.StatusOK()
}

// getStatusChange - get status change by assignee and class
func getStatusChange(status, assigneeClass string, assigneeOrigin core.Origin) ([]EntryStatusChange, *core.Status) {
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
