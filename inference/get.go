package inference

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	if values == nil {
		return nil, nil, core.StatusNotFound()
	}
	defer safe.Lock()()
	entries = originFilter(values, entryData)
	if len(entries) == 0 {
		return nil, nil, core.StatusNotFound()
	}
	return entries, nil, core.StatusOK()
}
