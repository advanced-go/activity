package assignment

import (
	"context"
	"errors"
	"github.com/advanced-go/activity/module"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
	"time"
)

func put[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, resource, template string, body []T, values url.Values, insert pgxsql.InsertFuncT[T], args ...any) (h2 http.Header, status *core.Status) {
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
	_, status = insert(ctx, h, resource, template, body, args...)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return

}

func testInsert[T pgxsql.Scanner[T]](_ context.Context, _ http.Header, resource, template string, entries []T, args ...any) (tag pgxsql.CommandTag, status *core.Status) {
	assigneeId := ""
	if len(args) > 0 {
		if s, ok := args[0].(string); ok {
			assigneeId = s
		}
	}
	switch p := any(&entries).(type) {
	case *[]Entry:
		status = insertEntry(*p, assigneeId)
	case *[]EntryDetail:
		status = insertDetail(core.Origin{}, (*p)[0])
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(entries))
	}
	if status.OK() {
		tag.RowsAffected = int64(len(entries))
	}
	return
}

func insertEntry(entries []Entry, assigneeId string) *core.Status {
	defer safeEntry.Lock()()
	for _, e := range entries {
		es := EntryStatus{EntryId: e.EntryId, StatusId: 0, AgentId: e.AgentId, CreatedTS: time.Time{}, Status: OpenStatus, AssigneeId: assigneeId}
		e.CreatedTS = time.Now().UTC()
		e.EntryId = entryData[len(entryData)-1].EntryId + 1
		insertStatus(e.Origin(), es)
	}
	return core.StatusOK()
}

func insertDetail(o core.Origin, detail EntryDetail) *core.Status {
	e, ok := index.LookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	detail.EntryId = e.EntryId
	detail.DetailId = detailData[len(detailData)-1].DetailId + 1
	detail.CreatedTS = time.Now().UTC()
	detailData = append(detailData, detail)
	return core.StatusOK()
}

func insertStatus(o core.Origin, es EntryStatus) *core.Status {
	e, ok := index.LookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	defer safeStatus.Lock()()

	es.EntryId = e.EntryId
	es.StatusId = statusData[len(statusData)-1].StatusId + 1
	es.CreatedTS = time.Now().UTC()
	statusData = append(statusData, es)
	return core.StatusOK()
}

func insertStatusChange(o core.Origin, change EntryStatusChange) *core.Status {
	e, ok := index.LookupEntry(o)
	if !ok {
		return core.StatusNotFound()
	}
	change.EntryId = e.EntryId
	change.ChangeId = changeData[len(changeData)-1].ChangeId + 1
	change.CreatedTS = time.Now().UTC()
	changeData = append(changeData, change)
	return core.StatusOK()
}
