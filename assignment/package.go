package assignment

import (
	"context"
	"errors"
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
	assignmentStatusUpdate = "assignment-status-update"

	entryPath             = "assignment/entry"
	entryDetailPath       = "assignment/detail"
	entryStatusPath       = "assignment/status"
	entryStatusUpdatePath = "assignment/status-update"
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
	case entryStatusUpdatePath:
		return GetT[EntryStatusUpdate](ctx, h, values)
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(entries))
	}
	return
}

// Constraints - get/put type constraints
type Constraints interface {
	Entry | EntryDetail | EntryStatus | EntryStatusUpdate
}

// GetT - resource typed GET
func GetT[T Constraints](ctx context.Context, h http.Header, values url.Values) (entries []T, h2 http.Header, status *core.Status) {
	switch p := any(&entries).(type) {
	case *[]Entry:
		*p, h2, status = get[core.Log, Entry](ctx, h, values, assignment, "", nil)
	case *[]EntryDetail:
		*p, h2, status = get[core.Log, EntryDetail](ctx, h, values, assignmentDetail, "", nil)
	case *[]EntryStatus:
		*p, h2, status = get[core.Log, EntryStatus](ctx, h, values, assignmentStatus, "", nil)
	case *[]EntryStatusUpdate:
		*p, h2, status = get[core.Log, EntryStatusUpdate](ctx, h, values, assignmentStatusUpdate, "", nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(entries))
	}
	return
}

// Put - resource PUT, with optional content override
func Put[T Constraints](r *http.Request, body []T) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]T](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return put[core.Log, T](r.Context(), core.AddRequestId(r.Header), body)
}
