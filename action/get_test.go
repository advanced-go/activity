package action

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
)

func ExampleGet() {
	q := "region=*"
	entries, _, status := get[core.Output, Entry](nil, nil, uri.BuildValues(q), actionResource, "", nil)
	fmt.Printf("test: get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, len(entries))

	q = "region=*&order=desc"
	entries, _, status = get[core.Output, Entry](nil, nil, uri.BuildValues(q), actionResource, "", nil)
	fmt.Printf("test: get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, entries)

	//Output:
	//test: get("region=*") -> [status:OK] [entries:4]
	//test: get("region=*&order=desc") -> [status:OK] [entries:[{4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com test 0 0 0   0 open} {3 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az3  www.host1.com test 0 0 0   0 open} {2 director-1 2024-06-10 09:00:35 +0000 UTC us-west-1 usw1-az2  www.host2.com test 0 0 0   0 open} {1 director-1 2024-06-10 09:00:35 +0000 UTC us-west-1 usw1-az1  www.host1.com test 0 0 0   0 open}]]

}
