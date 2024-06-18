package inference

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func put[E core.ErrorHandler](ctx context.Context, h http.Header, body []Entry) (http.Header, *core.Status) {
	if len(body) == 0 {
		return nil, core.StatusOK()
	}
	defer safe.Lock()()
	entryData = append(entryData, body...)
	return nil, core.StatusOK()
}
