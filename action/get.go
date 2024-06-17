package action

import (
	"context"
	"github.com/advanced-go/activity/common"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, values url.Values, resource, template string, query pgxsql.QueryFuncT[T]) (entries []T, h2 http.Header, status *core.Status) {
	//var e E
	return nil, nil, core.StatusOK()
}

func getEntryByStatus(ctx context.Context, h http.Header, o core.Origin, status string) ([]Entry, *core.Status) {
	e, ok := common.lookupEntry(o)
	if !ok {
		return nil, core.StatusNotFound()
	}
	entryList.Lock()
	defer entryList.Unlock()

	for i := len(statusData) - 1; i >= 0; i-- {
		if statusData[i].EntryId == e.EntryId && statusData[i].Status == status {
			return []Entry{e}, core.StatusOK()
		}
	}
	return nil, core.StatusNotFound()
}

func getEntry(ctx context.Context, h http.Header, values map[string][]string) ([]Entry, *core.Status) {
	return FilterT[Entry](values, entryData, validEntry)
}

func getStatus(ctx context.Context, h http.Header, values map[string][]string) ([]EntryStatus, *core.Status) {
	return FilterT[EntryStatus](values, statusData, validStatus)
}
