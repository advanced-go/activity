package customer1

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"net/http"
)

const (
	customerId = "customer"
	eventId    = "event"
)

type results struct {
	h       http.Header
	failure bool
	addr    []address
	events  []log
	status  *core.Status
	handler core.ErrorHandler
}

func newResults(h http.Header, handler core.ErrorHandler) *results {
	r := new(results)
	r.h = h
	r.handler = handler
	return r
}

func (r *results) handleError(status *core.Status) {
	r.handler.Handle(status.WithRequestId(r.h))
	r.status = status
	r.failure = true
}

func (r *results) onResponse(item httpx.RequestItem, resp *http.Response, status *core.Status) (proceed bool) {
	if !status.OK() {
		r.handleError(status)
		return false
	}
	var status1 *core.Status

	switch item.Id {
	case customerId:
		r.addr, status1 = json.New[[]address](resp.Body, resp.Header)
		if !status1.OK() {
			r.handleError(status1)
			return false
		}
	case eventId:
		r.events, status1 = json.New[[]log](resp.Body, resp.Header)
		if !status1.OK() {
			r.handleError(status1)
			return false
		}
	}
	return true
}

func (r *results) buildEntries() []Entry {

	return nil
}
