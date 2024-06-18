package action

import (
	"context"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

func put[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, resource, template string, body []T, insert pgxsql.InsertFuncT[T]) (http.Header, *core.Status) {
	if len(body) != 1 {
		return nil, core.StatusBadRequest()
	}
	switch p := any(&body).(type) {
	case *[]Entry:
		return nil, insertEntry(*p)
	default:
		return nil, core.StatusBadRequest()
	}
}

func insertEntry(entries []Entry) *core.Status {
	if len(entries) != 1 {
		return core.StatusBadRequest()
	}
	e := entries[0]
	defer safeEntry.Lock()()

	e.CreatedTS = time.Now().UTC()
	e.Status = OpenStatus
	e.EntryId = entryData[len(entryData)-1].EntryId + 1
	entryData = append(entryData, e)
	return core.StatusOK()
}
