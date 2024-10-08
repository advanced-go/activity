package customer1

import (
	"fmt"
	"github.com/advanced-go/activity/module"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func ExampleExchange() {
	var e core.Output

	h := make(http.Header)
	h.Add(core.XRequestId, "9876-543")
	h.Add(core.XRelatesTo, "123-456-7890")
	h.Add(core.XFrom, module.Authority)

	values := make(url.Values)
	values.Add(customerKey, "D001")

	ex := newExchange(h, e)
	ex.buildRequests(nil, h, activity1EgressPath, values)
	fmt.Printf("test: buildRequests() -> [status:%v] [req-id:%v]\n", ex.failure, h.Get(core.XRequestId))

	if ex.failure == nil {
		ex.do()
		entries, h := ex.buildResults()
		fmt.Printf("test: do() -> [status:%v] [entries:%v] [header:%v]\n", ex.failure, entries, h)
	}

	//Output:
	//test: buildRequests() -> [status:<nil>] [req-id:9876-543]
	//test: do() -> [status:<nil>] [entries:[{{D001 123 Main  Anytown OH 12345 before-email@hotmail.com} [{{us-west oregon dc1 www.test-host.com google-search 123456} 2024-06-03 18:29:16.0447249 +0000 UTC 100 egress GET  200 500 100 10 RL}]}]] [header:map[Content-Type:[application/json]]]
	
}
