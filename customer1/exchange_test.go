package customer1

import (
	"fmt"
	"github.com/advanced-go/activity/module"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

func ExampleExchange() {
	var e core.Output

	h := core.AddRequestId(nil)
	h.Add(core.XRelatesTo, "123-456-7890")
	h.Add(core.XFrom, module.Authority)

	values := make(url.Values)
	values.Add(customerKey, "D001")

	ex := newExchange(h, e, buildRequests)
	ex.buildRequests(nil, h, activity1EgressPath, values)
	fmt.Printf("test: buildRequests() -> [status:%v] [req-id:%v]\n", ex.failure, h.Get(core.XRequestId))

	if ex.failure == nil {
		ex.do()
		entries, h := ex.buildResults()
		fmt.Printf("test: do() -> [status:%v] [entries:%v] [header:%v]\n", ex.failure, entries, h)
	}

	//Output:
	//fail

}
