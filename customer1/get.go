package customer1

import (
	"context"
	"fmt"
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
	if h != nil && h.Get(uri.XContentLocationResolver) != "" {
		return h
	}
	return httpx.SetHeader(h, uri.XContentLocationResolver, testrsc.Customer1Entry)
}

func get[E core.ErrorHandler](ctx context.Context, h http.Header, resource string, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	var e E

	h2 = httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeText)
	if values == nil || len(values) == 0 {
		return nil, h2, core.StatusNotFound()
	}
	// Test only
	h = testOverride(h)
	results, status1 := exchange(ctx, h, resource, values)
	if !status1.OK() {
		e.Handle(status1.WithRequestId(h))
		return nil, h2, status1
	}
	entries, status = buildEntries(results)
	if !status.OK() {
		e.Handle(status.WithRequestId(h))
		return nil, h2, status
	}
	// Test only
	entries = filter(entries, values)
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	} else {
		h2 = httpx.SetHeader(h2, httpx.ContentType, httpx.ContentTypeJson)
	}
	return
}

func onResponse(resp *http.Response, status *core.Status) (failure, proceed bool) {
	//fmt.Printf("[req:%v]\n [resp:%v]\n [status:%v]\n", resp.Request, resp, status)
	fmt.Printf("[status:%v]\n", status)
	return !status.OK(), true
}

func exchange(ctx context.Context, h http.Header, resource string, values url.Values) ([]core.ExchangeResult, *core.Status) {
	reqs, status := buildRequests(ctx, h, resource, values)
	if !status.OK() {
		return nil, status
	}
	return httpx.MultiExchange(reqs, onResponse)
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
