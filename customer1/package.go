package customer1

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/customer/address1"
	Route   = "customer-activity"

	CustomerHost      = "localhost:8082"
	CustomerAuthority = "github/advanced-go/customer"
	CustomerPath      = "v1/address/entry"

	ObservationHost        = "localhost:8083"
	ObservationAuthority   = "github/advanced-go/observation"
	ObservationIngressPath = "v1/timeseries/ingress/entry"
	ObservationEgressPath  = "v1/timeseries/egress/entry"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)

// Get - customer1 resource GET
func Get(r *http.Request, _ string) (entries []Entry, h2 http.Header, status *core.Status) {
	if r == nil {
		return entries, h2, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
	}
	return nil, h2, core.StatusNotFound()
}
