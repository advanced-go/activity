package inference

import (
	"context"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
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

func insertEntry(body []Entry) *core.Status {
	if len(body) == 0 {
		return core.StatusOK()
	}
	defer safe.Lock()()
	entryData = append(entryData, body...)
	return core.StatusOK()
}
