package customer1

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

func buildResults(r *response) ([]Entry, http.Header, *core.Status) {

	// Build header
	h2 := httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeJson)

	// Build entries
	entry := Entry{
		Customer: r.addr.content[0],
		Activity: r.event.content,
	}
	return []Entry{entry}, h2, core.StatusOK()
}
