package customer1

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"net/http"
	"net/url"
)

func buildRequests(ctx context.Context, h http.Header, resource string, values url.Values) ([]httpx.RequestItem, *core.Status) {
	var reqs []httpx.RequestItem

	u := resolver.Url(CustomerHost, CustomerAuthority, CustomerV1AddressPath, values, h)
	req, err := http.NewRequestWithContext(core.NewContext(ctx), http.MethodGet, u, nil)
	if err != nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, err)
	}
	httpx.Forward(req.Header, h)
	reqs = append(reqs, httpx.RequestItem{Id: customerId, Request: req})

	switch resource {
	case logPath:
		u = resolver.Url(LogHost, LogAuthority, LogV1EventPath, values, h)
		req, err = http.NewRequestWithContext(core.NewContext(ctx), http.MethodGet, u, nil)
		if err != nil {
			return nil, core.NewStatusError(core.StatusInvalidArgument, err)
		}
		httpx.Forward(req.Header, h)
		reqs = append(reqs, httpx.RequestItem{Id: eventId, Request: req})
	default:
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: invalid resource %v", resource)))
	}
	return reqs, core.StatusOK()
}
