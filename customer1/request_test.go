package customer1

import (
	"fmt"
	"net/url"
)

func ExampleBuild() {
	values := make(url.Values)

	values.Add(customerKey, "C001")
	reqs, status := buildRequests(nil, nil, activity1IngressPath, values)

	fmt.Printf("test: buildRequests() -> [status:%v] [reqs:%v]\n", status, reqs[0])
	fmt.Printf("test: buildRequests() -> [status:%v] [reqs:%v]\n", status, reqs[1])

	//Output:
	//fail
}
