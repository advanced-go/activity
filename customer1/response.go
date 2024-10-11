package customer1

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
	"net/http"
	"sync"
)

const (
	customerId = "cust"
	eventId    = "event"
)

type result[T any] struct {
	content T
	resp    *http.Response
	status  *core.Status
}

func (r result[T]) StatusCode() int {
	if r.resp == nil {
		return http.StatusInternalServerError
	}
	return r.resp.StatusCode
}

func (r result[T]) Header() http.Header {
	if r.resp == nil {
		return make(http.Header)
	}
	return r.resp.Header
}

func (r result[T]) Status() *core.Status {
	return r.status
}

type response struct {
	nonSuccess []*core.Status
	mu         sync.RWMutex

	addr  result[[]address]
	event result[[]log]
}

func (r *response) isNonSuccessful() bool {
	return len(r.nonSuccess) > 0
}

func (r *response) addNonSuccessful(status *core.Status) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.nonSuccess = append(r.nonSuccess, status)
}

func (r *response) handler(id string, resp *http.Response, status *core.Status) {
	switch id {
	case customerId:
		r.addr.status = status
		r.addr.resp = resp
		if status.OK() {
			r.addr.content, r.addr.status = json.New[[]address](resp.Body, resp.Header)
		}
		if !r.addr.status.OK() {
			r.addNonSuccessful(r.addr.status)
		}
	case eventId:
		r.event.status = status
		r.event.resp = resp
		if status.OK() {
			r.event.content, r.event.status = json.New[[]log](resp.Body, resp.Header)
		}
		if !r.event.status.OK() {
			r.addNonSuccessful(r.event.status)
		}
	}
}

func verifyResponses[E core.ErrorHandler](r *response, h http.Header) *core.Status {
	var e E

	if !r.isNonSuccessful() {
		return core.StatusOK()
	}
	cnt := 0
	for _, status := range r.nonSuccess {
		if status.NotFound() {
			continue
		}
		cnt++
		e.Handle(status.WithRequestId(h))
	}
	if cnt == 0 {
		return core.StatusOK()
	}
	return core.NewStatusError(http.StatusInternalServerError, errors.New("multiple errors"))
}
