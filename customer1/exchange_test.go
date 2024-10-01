package customer1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func ExampleExchange() {
	var e core.Output

	h := make(http.Header)
	values := make(url.Values)
	values.Add(customerKey, "C001")

	ex := newExchange(h, e, buildRequests)
	ex.buildRequests(nil, h, activity1IngressPath, values)
	fmt.Printf("test: buildRequests() -> [status:%v]\n", ex.failure)

	if ex.failure == nil {
		ex.do()
		fmt.Printf("test: do() -> [status:%v]\n", ex.failure)
	}

	//Output:
	//fail

}
