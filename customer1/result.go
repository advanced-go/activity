package customer1

import (
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"net/http"
)

func buildResults(r *response) ([]Entry, http.Header, *core.Status) {

	// Build header
	h2 := httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeJson)
	if r.addr.Status().NotFound() {
		return []Entry{}, httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeText), r.addr.Status()
	}

	// Build entries
	entry := Entry{
		Customer: r.addr.content[0],
		Activity: r.event.content,
	}
	return []Entry{entry}, h2, core.StatusOK()
}
