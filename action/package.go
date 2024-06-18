package action

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
	PkgPath      = "github/advanced-go/activity/action"
	action       = "action"
	actionStatus = "action-status"

	entryPath       = "assignment/entry"
	entryStatusPath = "assignment/status"
)

// Get - resource GET
func Get(ctx context.Context, path string, h http.Header, values url.Values) (entries any, h2 http.Header, status *core.Status) {
	switch path {
	case entryPath:
		return GetT[Entry](ctx, h, values)
	case entryStatusPath:
		return GetT[EntryStatus](ctx, h, values)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: invalid path %v", path)))
	}
	return
}

// Constraints - get/put type constraints
type Constraints interface {
	Entry | EntryStatus
}

// GetT - typed resource GET
func GetT[T Constraints](ctx context.Context, h http.Header, values url.Values) (entries []T, h2 http.Header, status *core.Status) {
	switch p := any(&entries).(type) {
	case *[]Entry:
		*p, status = getEntry(ctx, h, values)
	case *[]EntryStatus:
		*p, status = getStatus(ctx, h, values)
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
	//case entryStatusPath:
	//	return PutT[EntryStatus](r, nil)
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
		status = insertEntry(r.Context(), core.AddRequestId(r.Header), *p)
		//	case *[]EntryStatus:
		//		status = insertStatus(r.Context(), core.AddRequestId(r.Header), *p)
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(body))
	}
	return
}

// GetEntryByStatus - by status
func GetEntryByStatus(ctx context.Context, h http.Header, o core.Origin, status string) ([]Entry, *core.Status) {
	return getEntryByStatus(ctx, h, o, status)
}

// InsertEntry - add entry
func InsertEntry(ctx context.Context, h http.Header, e Entry) *core.Status {
	return insertEntry[core.Log](ctx, h, []Entry{e})
}

// InsertStatus - add status
func InsertStatus(ctx context.Context, h http.Header, o core.Origin, es EntryStatus) *core.Status {
	return insertStatus[core.Log](ctx, h, o, []EntryStatus{es})
}
