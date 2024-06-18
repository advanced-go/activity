package action

import (
	"context"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

func put[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, resource, template string, body []T, insert pgxsql.InsertFuncT[T]) (h2 http.Header, status *core.Status) {
	//var e E
	return nil, core.StatusOK()
}

func insertEntry(ctx context.Context, h http.Header, entries []Entry) *core.Status {
	e := entries[0]
	es := EntryStatus{
		EntryId:   e.EntryId,
		StatusId:  0,
		AgentId:   e.AgentId,
		CreatedTS: time.Time{},
		Status:    OpenStatus,
	}
	defer safeEntry.Lock()()

	e.CreatedTS = time.Now().UTC()
	e.EntryId = entryData[len(entryData)-1].EntryId + 1
	index.AddEntry(e)
	entryData = append(entryData, e)
	return insertStatus(ctx, h, e.Origin(), []EntryStatus{es})
}

func insertStatus(ctx context.Context, h http.Header, o core.Origin, entries []EntryStatus) *core.Status {
	e, ok := index.LookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeStatus.Lock()()

	es := entries[0]
	es.EntryId = e.EntryId
	es.StatusId = statusData[len(statusData)-1].StatusId + 1
	es.CreatedTS = time.Now().UTC()
	statusData = append(statusData, es)
	return core.StatusOK()
}
