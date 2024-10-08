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
	ex := newExchange(h, e)
	ex.buildRequests(ctx, h, resource, values)
	if ex.failure != nil {
		return nil, h2, ex.failure
	}
	// Do multi exchange
	ex.do()
	if ex.failure != nil {
		return nil, h2, ex.failure
	}
	// Build results
	entries, h2 = ex.buildResults()
	if ex.failure != nil {
		return nil, h2, ex.failure
	}

	// Test only
	entries = filter(entries, values)
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	} else {
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
