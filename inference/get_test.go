package inference

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
)

func ExampleGet() {
	q := "region=*"
	entries, _, status := get[core.Output](nil, nil, uri.BuildValues(q))
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, len(entries))

	q = "region=*&order=desc"
	entries, _, status = get[core.Output](nil, nil, uri.BuildValues(q))
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, entries)

	//Output:
	//test: Get("region=*") -> [status:OK] [entries:2]
	//test: Get("region=*&order=desc") -> [status:OK] [entries:[{2 agent 2024-06-10 09:00:35 +0000 UTC us-west oregon  www.host2.com host text processed} {1 agent 2024-06-10 09:00:35 +0000 UTC us-west oregon  www.host1.com route information processed}]]

}
