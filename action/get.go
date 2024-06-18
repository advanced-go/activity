package action

import (
	"context"
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
	e, ok := index.LookupEntry(o)
	if !ok {
		return nil, core.StatusNotFound()
	}
	defer safeStatus.Lock()()

	for i := len(statusData) - 1; i >= 0; i-- {
		if statusData[i].EntryId == e.EntryId && statusData[i].Status == status {
			return []Entry{e}, core.StatusOK()
		}
	}
	return nil, core.StatusNotFound()
}

func getEntry(ctx context.Context, h http.Header, values map[string][]string) ([]Entry, *core.Status) {
	defer safeEntry.Lock()()
	entries, status := FilterT[Entry](values, entryData, validEntry)
	if !status.OK() {
		return nil, status
	}
	for i, e := range entries {
		entries[i].Status = findStatus(e.EntryId)
	}
	return entries, status
}

func getStatus(ctx context.Context, h http.Header, values map[string][]string) ([]EntryStatus, *core.Status) {
	defer safeStatus.Lock()()
	return FilterT[EntryStatus](values, statusData, validStatus)
}

func findStatus(entryId int) string {
	for i := len(statusData) - 1; i >= 0; i-- {
		if statusData[i].EntryId == entryId {
			return statusData[i].Status
		}
	}
	return ""
}
