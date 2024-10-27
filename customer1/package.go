package customer1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/jsonx"
	"github.com/advanced-go/common/uri"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/activity/customer1"
	Route   = "customer-activity"
	logPath = "customer/event/entry"
	//activityEgressPath  = "customer/egress/entry"

	CustomerHost          = "localhost:8082"
	CustomerAuthority     = "github/advanced-go/customer"
	CustomerV1AddressPath = "v1/address/entry"

	LogHost        = "localhost:8083"
	LogAuthority   = "github/advanced-go/log"
	LogV1EventPath = "v1/event/entry"
	//EventsV1EgressPath  = "v1/log/egress/entry"
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
	case logPath:
		t, h2, status := get[E](r.Context(), core.AddRequestId(r.Header), logPath, r.URL.Query())
		if !status.OK() {
			return nil, h2, status
		}
		buf, status1 := jsonx.Marshal(t)
		if !status1.OK() {
			e.Handle(status1)
			return nil, h2, status1
		}
		return buf, h2, status1
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: resource path is invalid [%v]", path)))
		return nil, nil, status
	}
}
