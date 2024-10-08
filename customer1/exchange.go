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

type requestsFunc func(ctx context.Context, h http.Header, resource string, values url.Values) ([]httpx.RequestItem, *core.Status)

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

	reqs     []httpx.RequestItem
	failure  *core.Status
	handler  core.ErrorHandler
	requests requestsFunc
}

func newExchange(h http.Header, handler core.ErrorHandler, requests requestsFunc) *exchange {
	r := new(exchange)
	r.h = h
	if r.h == nil {
		r.h = make(http.Header)
	}
	r.handler = handler
	r.requests = requests
	return r
}

func (e *exchange) handleError(status *core.Status) {
	e.handler.Handle(status.WithRequestId(e.h))
	e.failure = status
}

func (e *exchange) buildRequests(ctx context.Context, h http.Header, resource string, values url.Values) {
	e.reqs, e.failure = e.requests(ctx, h, resource, values)
	if !e.failure.OK() {
		e.handleError(e.failure)
	} else {
		e.failure = nil
	}
}

func (e *exchange) onResponse(id string, resp *http.Response, status *core.Status) (proceed bool) {
	// Check for connectivity errors, Gateway Timeout, Too Many Requests, and Internal Server Error
	// TODO: verify status for connectivity errors.
	switch status.Code {
	case http.StatusGatewayTimeout, http.StatusTooManyRequests, http.StatusInternalServerError:
		e.handleError(status)
		return false
	//case http.StatusNotFound:
	default:
		// If not OK then return as there is no content
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("status code : %v\n", resp.StatusCode)
			return true
		}
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
	//if !e.addr.OK() {
	//	e.handleError(e.addr.core.NewStatusError(core.StatusInvalidContent, errors.New("error: a response is not OK")).WithRequestId(e.h))
	//	return nil, nil
	//}
	if !e.addr.OK() || !e.event.OK() {
		e.handleError(core.NewStatusError(core.StatusInvalidContent, errors.New("error: a response is not OK")).WithRequestId(e.h))
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
