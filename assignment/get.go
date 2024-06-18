package assignment

import (
	"context"
	"fmt"
	"github.com/advanced-go/activity/common"
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
	o := NewOrigin(values)
	if o.Region != "*" {
		if e, ok := index.LookupEntry(o); ok {
			values["entry-id"] = []string{fmt.Sprintf("%v", e.EntryId)}
		}
	}
	switch p := any(&entries).(type) {
	case *[]Entry:
		defer safeEntry.Lock()()
		*p, status = common.FilterT[Entry](values, entryData, validEntry)
	case *[]EntryDetail:
		defer safeDetail.Lock()()
		*p, status = common.FilterT[EntryDetail](values, detailData, validDetail)
	case *[]EntryStatus:
		defer safeStatus.Lock()()
		*p, status = FilterT[EntryStatus](values, statusData, validStatus)
	case *[]EntryStatusChange:
		defer safeChange.Lock()()
		*p, status = FilterT[EntryStatusChange](values, changeData, validStatusChange)
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(entries))
	}
	return
}

func getStatusChange(ctx context.Context, h http.Header, values url.Values) ([]EntryStatusChange, *core.Status) {
	e, ok := index.LookupEntry(core.NewOrigin(values))
	if !ok {
		return nil, core.StatusNotFound()
	}
	defer safeChange.Lock()()
	cls := ""
	s := values["assignee-class"]
	if len(s) > 0 {
		cls = s[0]
	}
	for _, change := range changeData {
		if change.EntryId == e.EntryId && change.AssigneeClass == cls {
			return []EntryStatusChange{change}, core.StatusOK()
		}
	}
	return nil, core.StatusNotFound()
}

func NewOrigin(values map[string][]string) core.Origin {
	region := ""
	zone := ""
	subZone := ""
	host := ""

	s := values[core.RegionKey]
	if len(s) > 0 {
		region = s[0]
	}
	s = values[core.ZoneKey]
	if len(s) > 0 {
		zone = s[0]
	}
	s = values[core.SubZoneKey]
	if len(s) > 0 {
		subZone = s[0]
	}
	s = values[core.HostKey]
	if len(s) > 0 {
		host = s[0]
	}
	return core.Origin{Region: region, Zone: zone, SubZone: subZone, Host: host}
}
