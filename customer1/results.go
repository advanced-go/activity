package customer1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
)

func buildEntries(results []core.ExchangeResult) ([]Entry, *core.Status) {
	var e []Entry

	for _, r := range results {
		if r.Failure {
			return e, core.NewStatusError(core.StatusExecError, errors.New(fmt.Sprintf("error: failure on exchange request %v", r.Resp.Request.URL.String())))
		}
	}
	return e, core.StatusOK()

}
