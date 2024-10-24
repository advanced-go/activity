package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/activity/module"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"github.com/advanced-go/common/uri"
	"net/http"
)

const (
	PkgPath             = "github/advanced-go/activity/http"
	ver1                = "v1"
	ver2                = "v2"
	healthLivenessPath  = "health/liveness"
	healthReadinessPath = "health/readiness"
	versionPath         = "version"
	authorityPath       = "authority"
)

//var (
//	authorityResponse = NewAuthorityResponse(module.Authority)
//)

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if r == nil {
		status := core.NewStatusError(http.StatusBadRequest, errors.New("request is nil"))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
	p, err := uri.ValidateURL(r.URL, module.Authority)
	if err != nil {
		status := core.NewStatusError(http.StatusBadRequest, err)
		resp, _ := httpx.NewResponse(http.StatusBadRequest, h2, err)
		return resp, status
	}
	core.AddRequestId(r.Header)
	switch p.Resource {
	case module.CustomerResource:
		return customerExchange(r, p)
	case versionPath:
		return NewVersionResponse(module.Version), core.StatusOK()
	case authorityPath:
		return NewAuthorityResponse(module.Authority), core.StatusOK()
	case healthReadinessPath, healthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status := core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, resource not found: [%v]", p.Resource)))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
}
