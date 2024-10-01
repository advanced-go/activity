package customer1

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

func buildRequests(ctx context.Context, h http.Header, resource string, values url.Values) ([]httpx.RequestItem, *core.Status) {
	var reqs []httpx.RequestItem

	u := resolver.Url(CustomerHost, CustomerAuthority, Customer1AddressPath, values, h)
	req, err := http.NewRequestWithContext(core.NewContext(ctx), http.MethodGet, u, nil)
	if err != nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, err)
	}
	reqs = append(reqs, httpx.RequestItem{Id: customerId, Request: req})

	switch resource {
	case activity1IngressPath:
		u = resolver.Url(EventsHost, EventsAuthority, Events1IngressPath, values, h)
		req, err = http.NewRequestWithContext(core.NewContext(ctx), http.MethodGet, u, nil)
		if err != nil {
			return nil, core.NewStatusError(core.StatusInvalidArgument, err)
		}
		reqs = append(reqs, httpx.RequestItem{Id: eventId, Request: req})
	case activity1EgressPath:
		u = resolver.Url(EventsHost, EventsAuthority, Events1EgressPath, values, h)
		req, err = http.NewRequestWithContext(core.NewContext(ctx), http.MethodGet, u, nil)
		if err != nil {
			return nil, core.NewStatusError(core.StatusInvalidArgument, err)
		}
		reqs = append(reqs, httpx.RequestItem{Id: eventId, Request: req})
	default:
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: invalid resource %v", resource)))
	}
	return reqs, core.StatusOK()
}