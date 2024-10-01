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

type response[T any] struct {
	content T
	resp    *http.Response
}

func (r response[T]) StatusCode() int {
	if r.resp == nil {
		return http.StatusInternalServerError
	}
	return r.resp.StatusCode
}

func (r response[T]) Header() http.Header {
	if r.resp == nil {
		return make(http.Header)
	}
	return r.resp.Header
}

func (r response[T]) OK() bool {
	if r.resp == nil {
		return false
	}
	return r.StatusCode() == http.StatusOK
}

type exchange struct {
	h http.Header

	addr  response[[]address]
	event response[[]log]

	reqs    []httpx.RequestItem
	failure *core.Status
	handler core.ErrorHandler
}

func newExchange(h http.Header, handler core.ErrorHandler) *exchange {
	r := new(exchange)
	r.h = h
	r.handler = handler
	return r
}

func (e *exchange) handleError(status *core.Status) {
	e.handler.Handle(status.WithRequestId(e.h))
	e.failure = status
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

func (e *exchange) onResponse(id string, resp *http.Response, status *core.Status) (proceed bool) {
	// Check for connectivity errors, Gateway Timeout, Too Many Requests, and Internal Server Error
	// TODO: verify status for connectivity errors.
	if status.Code == http.StatusGatewayTimeout || status.Code == http.StatusTooManyRequests || status.Code == http.StatusInternalServerError {
		e.handleError(status)
		return false
	}
	var status1 *core.Status

	switch id {
	case customerId:
		e.addr.resp = resp
		e.addr.content, status1 = json.New[[]address](resp.Body, resp.Header)
		if !status1.OK() {
			e.handleError(status1)
			return false
		}
	case eventId:
		e.event.resp = resp
		e.event.content, status1 = json.New[[]log](resp.Body, resp.Header)
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

// TODO : verify responses, create entries, and set header values.
func (e *exchange) buildResults() ([]Entry, http.Header) {
	// Verify responses
	if !e.addr.OK() || !e.event.OK() {
		e.handleError(core.NewStatusError(core.StatusInvalidContent, errors.New("error: a response is mot OK")).WithRequestId(e.h))
		return nil, nil
	}
	// Build header
	h := httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeJson)

	// Build entries
	entry := Entry{
		Customer: e.addr.content[0],
		Activity: e.event.content,
	}
	return []Entry{entry}, h
}

/*

//resp    map[string]*http.Response
	//mu      sync.RWMutex

func (e *exchange) addResponse(id string, resp *http.Response) {
	if resp == nil || id == "" {
		return
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	e.resp[id] = resp
}

func (e *exchange) response(id string) *http.Response {
	if id == "" {
		e.handleError(core.NewStatusError(core.StatusInvalidArgument, errors.New("error: response id is empty")))
		return nil
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	resp := e.resp[id]
	if resp == nil {
		e.handleError(core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: response not found %v", id))))
	}
	return resp
}


*/
