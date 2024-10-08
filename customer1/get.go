package customer1

import (
	"context"
	"github.com/advanced-go/activity/testrsc"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

const (
	customerKey = "customer"
	stateKey    = "state"
)

func testOverride(h http.Header) http.Header {
	if h != nil && h.Get(uri.XContentResolver) != "" {
		return h
	}
	return httpx.SetHeader(h, uri.XContentResolver, testrsc.CustomerV1Entry)
}

func get[E core.ErrorHandler](ctx context.Context, h http.Header, resource string, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	var e E

	h2 = httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeText)
	if values == nil || len(values) == 0 {
		return nil, h2, core.StatusNotFound()
	}
	// Test only
	h = testOverride(h)

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

	// Test only
	entries = filter(entries, values)
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	} else {
		h2 = httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeJson)
		status = core.StatusOK()
	}
	return
}

func filter(entries []Entry, values url.Values) (result []Entry) {
	customer := values.Get(customerKey)
	state := values.Get(stateKey)
	for _, e := range entries {
		if customer == "*" {
			result = append(result, e)
			continue
		}
		if customer != "" && customer != e.CustomerId() {
			continue
		}
		if state != "" && state != e.State() {
			continue
		}
		result = append(result, e)
	}
	return result
}
