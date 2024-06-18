package assignment

import (
	"errors"
	"fmt"
	"github.com/advanced-go/activity/common"
	"net/url"
	"time"
)

var (
	safeChange = common.NewSafe()
	changeData = []EntryStatusChange{
		{EntryId: 1, ChangeId: 1, AgentId: "agent-name:agent-class:instance-id", AssigneeClass: "class", NewStatus: "closed", NewAssigneeClass: "new", Error: "test error", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 1, ChangeId: 2, AgentId: "agent-name:agent-class:instance-id", AssigneeClass: "class2", NewStatus: "closed", NewAssigneeClass: "new", Error: "test2 error", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// EntryStatusChange - updates for reassignment and close
type EntryStatusChange struct {
	EntryId       int       `json:"entry-id"`
	ChangeId      int       `json:"change-id"`
	AgentId       string    `json:"agent-id"` // Creation agent id
	CreatedTS     time.Time `json:"created-ts"`
	AssigneeClass string    `json:"assignee-class"` // Used to determine which class of agents will receive this change

	// Update data. Processed agent id needed ??
	// Error needed if updates are in an invalid order, such as a reassignment after a close
	NewStatus        string    `json:"new-status"`         // Status can be 'closed' or 'reassignment'
	NewAssigneeClass string    `json:"new-assignee-class"` // On reassignment, set new owner
	Error            string    `json:"error"`
	UpdatedTS        time.Time `json:"updated-ts"`
}

func (EntryStatusChange) Scan(columnNames []string, values []any) (e EntryStatusChange, err error) {
	for i, name := range columnNames {
		switch name {
		case EntryIdName:
			e.EntryId = values[i].(int)
		case ChangeIdName:
			e.ChangeId = values[i].(int)
		case AgentIdName:
			e.AgentId = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)
		case AssigneeClassName:
			e.AssigneeClass = values[i].(string)
		case NewStatusName:
			e.NewStatus = values[i].(string)
		case NewAssigneeClassName:
			e.NewAssigneeClass = values[i].(string)
		case ErrorName:
			e.Error = values[i].(string)
		case ProcessedTSName:
			e.UpdatedTS = values[i].(time.Time)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (a EntryStatusChange) Values() []any {
	return []any{
		a.EntryId,
		a.ChangeId,
		a.AgentId,
		a.CreatedTS,
		a.AssigneeClass,
		a.NewStatus,
		a.NewAssigneeClass,
		a.Error,
		a.UpdatedTS,
	}
}

func (EntryStatusChange) Rows(entries []EntryStatusChange) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}

func validStatusChange(values url.Values, e EntryStatusChange) bool {
	if values == nil || values.Get("entry-id") != fmt.Sprintf("%v", e.EntryId) {
		return false
	}

	// Additional filtering
	return true
}
