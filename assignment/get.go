package assignment

import (
	"context"
	"github.com/advanced-go/activity/module"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, values url.Values, resource, template string, query pgxsql.QueryFuncT[T]) (entries []T, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, h2, core.StatusNotFound()
	}
	if query == nil {
		query = pgxsql.QueryT[T] //testQuery[T] //pgxsql.QueryT[T]
	}
	h2 = httpx.Forward(h2, h)
	h2.Set(core.XFrom, module.Authority)
	entries, status = query(ctx, h, resource, template, values)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return
}

func testQuery[T pgxsql.Scanner[T]](_ context.Context, _ http.Header, _, _ string, values map[string][]string, _ ...any) (entries []T, status *core.Status) {
	switch p := any(&entries).(type) {
	case *[]Entry:
		*p, status = FilterT[Entry](values, entryData, validEntry)
	case *[]EntryDetail:
		*p, status = FilterT[EntryDetail](values, detailData, validDetail)
	case *[]EntryStatus:
		*p, status = FilterT[EntryStatus](values, statusData, validStatus)
	case *[]EntryStatusChange:
		*p, status = FilterT[EntryStatusChange](values, changeData, validStatusUpdate)
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(entries))
	}
	return
}

func getEntryByStatus(cx context.Context, h http.Header, o core.Origin, status string) ([]Entry, *core.Status) {
	e, ok := lookupEntry(o)
	if !ok {
		return nil, core.StatusNotFound()
	}
	for i := len(statusData) - 1; i >= 0; i-- {
		if statusData[i].EntryId == e.EntryId && statusData[i].Status == status {
			return []Entry{e}, core.StatusOK()
		}
	}
	return nil, core.StatusNotFound()
}

func getStatusChange[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, assigneeClass string) ([]EntryStatusChange, *core.Status) {
	e, ok := lookupEntry(o)
	if !ok {
		return nil, core.StatusNotFound()
	}
	for _, change := range changeData {
		if change.EntryId == e.EntryId && change.AssigneeClass == assigneeClass {
			return []EntryStatusChange{change}, core.StatusOK()
		}
	}
	return nil, core.StatusNotFound()
}
