package assignment

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

const (
	PkgPath                = "github/advanced-go/activity/assignment"
	assignment             = "assignment"
	assignmentDetail       = "assignment-detail"
	assignmentStatus       = "assignment-status"
	assignmentStatusChange = "assignment-status-change"

	assignmentPath             = "assignment"
	assignmentDetailPath       = "assignment/detail"
	assignmentStatusPath       = "assignment/status"
	assignmentStatusChangePath = "assignment/status-change"
)

// Get - resource GET
func Get(ctx context.Context, path string, h http.Header, values url.Values) (entries any, h2 http.Header, status *core.Status) {
	switch path {
	case assignmentPath:
		return GetT[Entry](ctx, h, values)
	case assignmentDetailPath:
		return GetT[EntryDetail](ctx, h, values)
	case assignmentStatusPath:
		return GetT[EntryStatus](ctx, h, values)
	case assignmentStatusChangePath:
		return GetT[EntryStatusChange](ctx, h, values)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid path %v", path)))
	}
	return
}

// Constraints - get/put type constraints
type Constraints interface {
	Entry | EntryDetail | EntryStatus | EntryStatusChange
}

// GetT - typed resource GET
func GetT[T Constraints](ctx context.Context, h http.Header, values url.Values) (entries []T, h2 http.Header, status *core.Status) {
	switch p := any(&entries).(type) {
	case *[]Entry:
		*p, h2, status = get[core.Log, Entry](ctx, h, values, assignment, "", nil)
	case *[]EntryDetail:
		*p, h2, status = get[core.Log, EntryDetail](ctx, h, values, assignmentDetail, "", nil)
	case *[]EntryStatus:
		*p, h2, status = get[core.Log, EntryStatus](ctx, h, values, assignmentStatus, "", nil)
	case *[]EntryStatusChange:
		*p, h2, status = get[core.Log, EntryStatusChange](ctx, h, values, assignmentStatusChange, "", nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(entries))
	}
	return
}

// Put - resource PUT
/*
func Put(r *http.Request, path string) (h2 http.Header, status *core.Status) {
	switch path {
	case entryPath:
		return PutT[Entry](r, nil)
	case entryDetailPath:
		return PutT[EntryDetail](r, nil)
	case entryStatusChangePath:
		return PutT[EntryStatusChange](r, nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid path %v", path)))
	}
	return
}

// PutT - typed resource PUT, with optional content override
func PutT[T Constraints](r *http.Request, body []T) (h2 http.Header, status *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status1 := json2.New[[]T](r.Body, r.Header)
		if !status1.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return nil, status1
		}
		body = content
	}

		switch p := any(&body).(type) {
		case *[]Entry:
			h2, status = put[core.Log, Entry](r.Context(), core.AddRequestId(r.Header), assignment, "", *p, nil)
		case *[]EntryDetail:
			h2, status = put[core.Log, EntryDetail](r.Context(), core.AddRequestId(r.Header), assignmentDetail, "", *p, nil)
		case *[]EntryStatusChange:
			h2, status = put[core.Log, EntryStatusChange](r.Context(), core.AddRequestId(r.Header), assignmentStatusChange, "", *p, nil)
		default:
			status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(body))
		}
	return
}

*/

// Insert - add an assignment and an open status
func Insert(ctx context.Context, agentId string, origin core.Origin, assigneeClass string, assigneeOrigin core.Origin) *core.Status {
	return insert(ctx, agentId, origin, assigneeClass, assigneeOrigin)
}

// GetOpen - find an open assignment for a given assignee class and origin
func GetOpen(ctx context.Context, assigneeClass string, assigneeOrigin core.Origin) ([]Entry, *core.Status) {
	return getOpen(ctx, assigneeClass, assigneeOrigin)
}

// Assign - set the status of an assignment to assigned
func Assign(ctx context.Context, origin core.Origin, agentId, assigneeId string) *core.Status {
	return assign(ctx, origin, agentId, assignmentStatus)
}

// Close - update the status of an assignment to closed
func Close(ctx context.Context, origin core.Origin, agentId string) *core.Status {
	return closeAssignment(ctx, origin, agentId)
}

// AddDetail - add assignment details
func AddDetail(ctx context.Context, origin core.Origin, agentId, routeName, details string) *core.Status {
	return addDetail(ctx, origin, agentId, routeName, details)
}

// GetCloseStatusChange - get status change by assignee for close status
func GetCloseStatusChange(ctx context.Context, assigneeClass string, assigneeOrigin core.Origin) ([]EntryStatusChange, *core.Status) {
	return getStatusChange(ctx, ClosedStatus, assigneeClass, assigneeOrigin)
}

// GetReassignmentStatusChange - get status change by assignee for reassignment status
func GetReassignmentStatusChange(ctx context.Context, assigneeClass string, assigneeOrigin core.Origin) ([]EntryStatusChange, *core.Status) {
	return getStatusChange(ctx, ReassignmentStatus, assigneeClass, assigneeOrigin)
}

// ProcessReassignment - process the reassignment
func ProcessReassignment(ctx context.Context, change []EntryStatusChange) *core.Status {
	return processReassignment(ctx, change)
}

// Reassign - set the status of an assignment to reassignment, and update assignment receiver
func Reassign(ctx context.Context, origin core.Origin, agentId, newAssigneeClass string, newAssigneeOrigin core.Origin) *core.Status {
	return reassign(ctx, origin, agentId, newAssigneeClass, newAssigneeOrigin)
}
