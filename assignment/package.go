package assignment

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	PkgPath                = "github/advanced-go/activity/assignment"
	assignment             = "assignment"
	assignmentDetail       = "assignment-detail"
	assignmentStatus       = "assignment-status"
	assignmentStatusChange = "assignment-status-change"

	entryPath             = "assignment/entry"
	entryDetailPath       = "assignment/detail"
	entryStatusPath       = "assignment/status"
	entryStatusChangePath = "assignment/status-change"
)

// Get - resource GET
func Get(ctx context.Context, path string, h http.Header, values url.Values) (entries any, h2 http.Header, status *core.Status) {
	switch path {
	case entryPath:
		return GetT[Entry](ctx, h, values)
	case entryDetailPath:
		return GetT[EntryDetail](ctx, h, values)
	case entryStatusPath:
		return GetT[EntryStatus](ctx, h, values)
	case entryStatusChangePath:
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

// GetStatusChange - get change
func GetStatusChange(ctx context.Context, h http.Header, values url.Values) ([]EntryStatusChange, *core.Status) {
	return getStatusChange(ctx, h, values)
}

// GetEntryByStatus - by status
func GetEntryByStatus(ctx context.Context, h http.Header, o core.Origin, status string) ([]Entry, *core.Status) {
	return nil, core.StatusOK() //getEntryByStatus(ctx, h, o, status)
}

// InsertEntry - add entry
func InsertEntry(ctx context.Context, entry Entry, assigneeId string) *core.Status {
	return insertEntry([]Entry{entry}, assigneeId)
}

// InsertDetail - add detail
func InsertDetail(ctx context.Context, o core.Origin, detail EntryDetail) *core.Status {
	return insertDetail(o, detail)
}

// InsertStatus - add status
func InsertStatus(ctx context.Context, o core.Origin, status EntryStatus) *core.Status {
	return insertStatus(o, status)
}

// InsertStatusChange - add status change
func InsertStatusChange(ctx context.Context, o core.Origin, change EntryStatusChange) *core.Status {
	return insertStatusChange(o, change)
}

// ReassignEntry - reassign
func ReassignEntry(ctx context.Context, h http.Header, o core.Origin, assigneeClass string) *core.Status {
	return reassignEntry[core.Log](ctx, h, o, assigneeClass)
}

// AssignEntry - assign an entry
func AssignEntry(ctx context.Context, h http.Header, o core.Origin, assigneeId string) *core.Status {
	return assignEntry[core.Log](ctx, h, o, assigneeId)
}
