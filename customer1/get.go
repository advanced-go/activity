package customer1

import (
	"context"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"net/http"
	"net/url"
)

const (
	customerKey = "customer"
)

func get[E core.ErrorHandler](ctx context.Context, h http.Header, resource string, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	var e E

	h2 = httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeText)
	if values == nil || len(values) == 0 {
		return nil, h2, core.StatusNotFound()
	}
	// Build requests
	reqs, status1 := buildRequests(ctx, h, resource, values)
	if !status1.OK() {
		e.Handle(status1.WithRequestId(h))
		return nil, h2, status1
	}

	// Create response and process exchanges
	resp := new(response)
	httpx.MultiExchange(reqs, resp.handler)

	// Verify responses
	status = verifyResponses[E](resp, h)
	if !status.OK() {
		return
	}

	// Build results
	entries, h2, status = buildResults(resp)
	if !status.OK() {
		e.Handle(status.WithRequestId(h))
		return
	}
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	} else {
		h2 = httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeJson)
		status = core.StatusOK()
	}
	return
}
