package action

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	PkgPath        = "github/advanced-go/activity/action"
	actionResource = "action"
)

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	return get[core.Log, Entry](ctx, h, values, actionResource, "", nil)
}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) (h2 http.Header, status *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status1 := json2.New[[]Entry](r.Body, r.Header)
		if !status1.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return nil, status1
		}
		body = content
	}
	return put[core.Log](r.Context(), core.AddRequestId(r.Header), actionResource, "", body, nil)
}

// Insert - add entry
func Insert(ctx context.Context, h http.Header, e Entry) *core.Status {
	_, status := put[core.Log, Entry](ctx, core.AddRequestId(h), actionResource, "", []Entry{e}, nil)
	return status
}

// Need to add an get entry for entries that are not open. Also needs to update status to closed
