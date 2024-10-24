package http

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"net/http"
)

const (
	versionFmt   = "{\n \"version\": \"%v\"\n}"
	authorityFmt = "{\n \"authority\": \"%v\"\n}"
)

func NewVersionResponse(version string) *http.Response {
	content := fmt.Sprintf(versionFmt, version)
	resp, _ := httpx.NewResponse(http.StatusOK, httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeText), content)
	return resp
}

func NewAuthorityResponse(authority string) *http.Response {
	content := fmt.Sprintf(authorityFmt, authority)
	resp, _ := httpx.NewResponse(http.StatusOK, httpx.SetHeader(nil, core.XAuthority, authority), content)
	return resp
}
