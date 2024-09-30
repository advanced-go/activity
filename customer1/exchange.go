package customer1

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	customerId = "customer"
	eventId    = "event"
)

type exchange struct {
	failure bool
	h       http.Header
	addr    []address
	events  []log
	reqs    []httpx.RequestItem
	status  *core.Status
	handler core.ErrorHandler
}

func newExchange(h http.Header, handler core.ErrorHandler) *exchange {
	r := new(exchange)
	r.h = h
	r.handler = handler
	return r
}

func (e *exchange) handleError(status *core.Status) *core.Status {
	e.handler.Handle(status.WithRequestId(e.h))
	e.status = status
	e.failure = true
	return status
}

func (e *exchange) buildRequests(ctx context.Context, h http.Header, resource string, values url.Values) {
	u := resolver.Url(CustomerHost, CustomerAuthority, Customer1AddressPath, values, h)
	req, err := http.NewRequestWithContext(core.NewContext(ctx), http.MethodGet, u, nil)
	if err != nil {
		e.handleError(core.NewStatusError(core.StatusInvalidArgument, err))
		return
	}
	e.reqs = append(e.reqs, httpx.RequestItem{Id: customerId, Request: req})

	switch resource {
	case activity1IngressPath:
		u = resolver.Url(EventsHost, EventsAuthority, Events1IngressPath, values, h)
		req, err = http.NewRequestWithContext(core.NewContext(ctx), http.MethodGet, u, nil)
		if err != nil {
			e.handleError(core.NewStatusError(core.StatusInvalidArgument, err))
			return
		}
		e.reqs = append(e.reqs, httpx.RequestItem{Id: eventId, Request: req})
	case activity1EgressPath:
		u = resolver.Url(EventsHost, EventsAuthority, Events1EgressPath, values, h)
		req, err = http.NewRequestWithContext(core.NewContext(ctx), http.MethodGet, u, nil)
		if err != nil {
			e.handleError(core.NewStatusError(core.StatusInvalidArgument, err))
			return
		}
		e.reqs = append(e.reqs, httpx.RequestItem{Id: eventId, Request: req})
	default:
		e.handleError(core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: invalid resource %v", resource))))
		return
	}
}

func (e *exchange) onResponse(item httpx.RequestItem, resp *http.Response, status *core.Status) (proceed bool) {
	if !status.OK() {
		e.handleError(status)
		return false
	}
	var status1 *core.Status

	switch item.Id {
	case customerId:
		e.addr, status1 = json.New[[]address](resp.Body, resp.Header)
		if !status1.OK() {
			e.handleError(status1)
			return false
		}
	case eventId:
		e.events, status1 = json.New[[]log](resp.Body, resp.Header)
		if !status1.OK() {
			e.handleError(status1)
			return false
		}
	}
	return true
}

func (e *exchange) do() {
	httpx.MultiExchange(e.reqs, e.onResponse)
}

func (e *exchange) buildEntries() []Entry {
	return nil
}
