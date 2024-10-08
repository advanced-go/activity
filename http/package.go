package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/activity/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/activity/http"
	ver1    = "v1"
	ver2    = "v2"
)

var (
	authorityResponse = httpx.NewAuthorityResponse(module.Authority)
)

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if r == nil {
		status := core.NewStatusError(http.StatusBadRequest, errors.New("request is nil"))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
	p, status := httpx.ValidateURL(r.URL, module.Authority)
	if !status.OK() {
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
	core.AddRequestId(r.Header)
	switch p.Resource {
	case module.CustomerResource:
		return customerExchange(r, p)
	case core.VersionPath:
		return httpx.NewVersionResponse(module.Version), core.StatusOK()
	case core.AuthorityPath:
		return authorityResponse, core.StatusOK()
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", p.Resource)))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
}
