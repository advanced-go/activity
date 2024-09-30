package customer1

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

const (
	PkgPath              = "github/advanced-go/activity/customer1"
	Route                = "customer-activity"
	activity1IngressPath = "v1/ingress/entry"
	activity1EgressPath  = "v1/egress/entry"

	CustomerHost         = "localhost:8082"
	CustomerAuthority    = "github/advanced-go/customer"
	Customer1AddressPath = "v1/address/entry"

	EventsHost         = "localhost:8083"
	EventsAuthority    = "github/advanced-go/events"
	Events1IngressPath = "v1/log/ingress/entry"
	Events1EgressPath  = "v1/log/egress/entry"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)

// Get - customer1 resource GET
func Get(r *http.Request, path string) ([]byte, http.Header, *core.Status) {
	if r == nil {
		return nil, nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
	}
	if r.Header.Get(core.XFrom) == "" {
		return httpGet[core.Log](r, path)
	}
	return httpGet[core.Output](r, path)
}

func httpGet[E core.ErrorHandler](r *http.Request, path string) ([]byte, http.Header, *core.Status) {
	var e E

	switch path {
	case activity1IngressPath:
		t, h2, status := get[E](r.Context(), core.AddRequestId(r.Header), activity1IngressPath, r.URL.Query())
		if !status.OK() {
			return nil, h2, status
		}
		buf, status1 := json2.Marshal(t)
		if !status1.OK() {
			e.Handle(status1)
			return nil, h2, status1
		}
		return buf, h2, status1
	case activity1EgressPath:
		t, h2, status := get[E](r.Context(), core.AddRequestId(r.Header), activity1EgressPath, r.URL.Query())
		if !status.OK() {
			return nil, h2, status
		}
		buf, status1 := json2.Marshal(t)
		if !status1.OK() {
			e.Handle(status1)
			return nil, h2, status1
		}
		return buf, h2, status1
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New("error: resource is not ingress or egress"))
		return nil, nil, status
	}
}
