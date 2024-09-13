package customer1

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

const (
	PkgPath      = "github/advanced-go/customer/address1"
	UpstreamPath = "storage/address"
	CustomerKey  = "customer"
	StateKey     = "state"
	Route        = "customer-address"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)

// AddressStorage - egress URLs
func AddressStorage(host, path string, values url.Values, h http.Header) string {
	return resolver.Url(host, path, values, h)
}

// Get - activity1 resource GET
func Get(r *http.Request, _ string) (entries []Entry, h2 http.Header, status *core.Status) {
	if r == nil {
		return entries, h2, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
	}
	return nil, h2, core.StatusNotFound()
}
