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

var index = make(map[string]string)

func init() {
	for _, e := range entryData {
		index[createKey(core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host})] = ""
	}
}

func get[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, values url.Values, resource, template string, query pgxsql.QueryFuncT[T]) (entries []T, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, h2, core.StatusNotFound()
	}
	if query == nil {
		query = testQuery[T] //pgxsql.QueryT[T]
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
		*p, status = FilterT[EntryDetail](values, entryDetailData, validDetail)
	case *[]EntryStatus:
		*p, status = FilterT[EntryStatus](values, entryStatusData, validStatus)
	case *[]EntryStatusUpdate:
		*p, status = FilterT[EntryStatusUpdate](values, entryStatusUpdateData, validStatusUpdate)
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(entries))
	}
	return
}

func validEntry(values url.Values, e Entry) bool {
	if values == nil {
		return false
	}
	filter := core.NewOrigin(values)
	target := core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}
	if !core.OriginMatch(target, filter) {
		return false
	}
	// Additional filtering
	return true
}

func validDetail(values url.Values, e EntryDetail) bool {
	if values == nil {
		return false
	}
	if !entryExists(values) {
		return false
	}
	// Additional filtering
	return true
}

func validStatus(values url.Values, e EntryStatus) bool {
	if values == nil {
		return false
	}
	if !entryExists(values) {
		return false
	}
	// Additional filtering
	return true
}

func validStatusUpdate(values url.Values, e EntryStatusUpdate) bool {
	if values == nil {
		return false
	}
	if !entryExists(values) {
		return false
	}
	// Additional filtering
	return true
}

func selectEntries(values url.Values) ([]Entry, *core.Status) {
	var entries []Entry

	for _, entry := range entryData {
		if validEntry(values, entry) {
			entries = append(entries, entry)
		}
	}
	if len(entries) == 0 {
		return nil, core.StatusNotFound()
	}
	return entries, core.StatusOK()
}

func createKey(o core.Origin) string {
	key := o.Region + ":"
	key += o.Zone + ":"
	key += o.Host
	return key
}

func entryExists(values url.Values) bool {
	_, ok := index[createKey(core.NewOrigin(values))]
	return ok
}
