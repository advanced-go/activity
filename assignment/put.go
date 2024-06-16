package assignment

import (
	"context"
	"errors"
	"github.com/advanced-go/activity/module"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"time"
)

func put[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, resource, template string, body []T, insert pgxsql.InsertFuncT[T]) (h2 http.Header, status *core.Status) {
	var e E

	if len(body) == 0 {
		status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: no entries found"))
		return nil, status
	}
	if insert == nil {
		insert = testInsert[T] //pgxsql.InsertT[T]
	}
	h2 = httpx.Forward(h2, h)
	h2.Set(core.XFrom, module.Authority)
	_, status = insert(ctx, h, resource, template, body)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return

}

func testInsert[T pgxsql.Scanner[T]](_ context.Context, _ http.Header, resource, template string, entries []T, args ...any) (pgxsql.CommandTag, *core.Status) {
	return pgxsql.CommandTag{}, core.NewStatus(http.StatusTeapot)
}

func insertEntry[E core.ErrorHandler](ctx context.Context, h http.Header, e Entry, assigneeId string) *core.Status {
	es := EntryStatus{
		EntryId:    e.EntryId,
		StatusId:   0,
		AgentId:    e.AgentId,
		CreatedTS:  time.Time{},
		Status:     OpenStatus,
		AssigneeId: assigneeId,
	}
	e.CreatedTS = time.Now().UTC()
	_, status := put[E, Entry](ctx, h, "", "", []Entry{e}, nil)
	if status.OK() {
		status = insertStatus[E](ctx, h, e.Origin(), es)
	}
	return status
}

func insertDetail[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, detail EntryDetail) *core.Status {
	e, ok := lookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	detail.EntryId = e.EntryId
	detail.DetailId = detailData[len(detailData)-1].DetailId + 1
	detail.CreatedTS = time.Now().UTC()
	_, status := put[E, EntryDetail](ctx, h, "", "", []EntryDetail{detail}, nil)
	return status
}

func insertStatus[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, es EntryStatus) *core.Status {
	e, ok := lookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	es.EntryId = e.EntryId
	es.StatusId = statusData[len(statusData)-1].StatusId + 1
	es.CreatedTS = time.Now().UTC()
	_, status := put[E, EntryStatus](ctx, h, "", "", []EntryStatus{es}, nil)
	return status
}

func insertStatusChange[E core.ErrorHandler](ctx context.Context, h http.Header, o core.Origin, change EntryStatusChange) *core.Status {
	e, ok := lookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	change.EntryId = e.EntryId
	change.ChangeId = changeData[len(changeData)-1].ChangeId + 1
	change.CreatedTS = time.Now().UTC()
	_, status := put[E, EntryStatusChange](ctx, h, "", "", []EntryStatusChange{change}, nil)
	return status
}
