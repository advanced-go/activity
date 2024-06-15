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
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid path %v", path)))
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

// Put - resource PUT
func Put(r *http.Request, path string) (h2 http.Header, status *core.Status) {
	switch path {
	case entryPath:
		return PutT[Entry](r, nil)
	case entryDetailPath:
		return PutT[EntryDetail](r, nil)
	case entryStatusPath:
		return PutT[EntryStatus](r, nil)
	case entryStatusUpdatePath:
		return PutT[EntryStatusUpdate](r, nil)
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
	case *[]EntryStatus:
		h2, status = put[core.Log, EntryStatus](r.Context(), core.AddRequestId(r.Header), assignmentStatus, "", *p, nil)
	case *[]EntryStatusUpdate:
		h2, status = put[core.Log, EntryStatusUpdate](r.Context(), core.AddRequestId(r.Header), assignmentStatusUpdate, "", *p, nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(body))
	}
	return
}
