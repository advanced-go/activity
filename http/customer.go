package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/activity/customer1"
	"github.com/advanced-go/activity/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

func customerExchange(r *http.Request, p *uri.Parsed) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if p == nil {
		p1, status := httpx.ValidateURL(r.URL, module.Authority)
		if !status.OK() {
			return httpx.NewResponse(status.HttpCode(), h2, status.Err)
		}
		p = p1
	}
	switch r.Method {
	case http.MethodGet:
		return customerGet(r, p)
	case http.MethodPut:
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
	return nil, core.StatusOK()
}

func customerGet(r *http.Request, p *uri.Parsed) (resp *http.Response, status *core.Status) {
	var entries any
	var h2 http.Header

	switch p.Version {
	case module.Ver1, "":
		entries, h2, status = customer1.Get(r, p.Path)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
	}
	if !status.OK() {
		resp, _ = httpx.NewResponse(status.HttpCode(), h2, status.Err)
		return resp, status
	}
	return httpx.NewResponse(status.HttpCode(), h2, entries)
}
