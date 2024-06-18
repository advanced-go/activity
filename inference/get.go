package inference

import (
	"context"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, values url.Values, resource, template string, query pgxsql.QueryFuncT[T]) (entries []T, h2 http.Header, status *core.Status) {
	switch p := any(&entries).(type) {
	case *[]Entry:
		*p, status = getEntry(values)
		return
	default:
		return nil, nil, core.StatusBadRequest()
	}
}

func getEntry(values url.Values) (entries []Entry, status *core.Status) {
	if values == nil {
		return nil, core.StatusNotFound()
	}
	defer safe.Lock()()
	entries = originFilter(values, entryData)
	if len(entries) == 0 {
		return nil, core.StatusNotFound()
	}
	return entries, core.StatusOK()
}
