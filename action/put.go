package action

import (
	"context"
	"github.com/advanced-go/activity/common"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

func put[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, resource, template string, body []T, insert pgxsql.InsertFuncT[T]) (h2 http.Header, status *core.Status) {
	//var e E
	return nil, core.StatusOK()
}

func insertEntry[E core.ErrorHandler](ctx context.Context, h http.Header, e Entry) *core.Status {
	es := EntryStatus{
		EntryId:   e.EntryId,
		StatusId:  0,
		AgentId:   e.AgentId,
		CreatedTS: time.Time{},
		Status:    OpenStatus,
	}
	entryList.Lock()
	defer entryList.Unlock()

	e.CreatedTS = time.Now().UTC()
	e.EntryId = entryData[len(entryData)-1].EntryId + 1
	common.index[e.Origin().Tag()] = e
	entryList.Append([]Entry{e})
	return insertStatus[E](ctx, h, e.Origin(), es)
}

func insertStatus[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, es EntryStatus) *core.Status {
	e, ok := common.lookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	statusList.Lock()
	defer statusList.Unlock()

	es.EntryId = e.EntryId
	es.StatusId = statusData[len(statusData)-1].StatusId + 1
	es.CreatedTS = time.Now().UTC()
	statusList.Append([]EntryStatus{es})
	return core.StatusOK()
}
