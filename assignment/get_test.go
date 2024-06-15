package assignment

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
)

func ExampleGet_Entry() {
	q := "region=*"
	entries, _, status := get[core.Output, Entry](nil, nil, uri.BuildValues(q), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, len(entries))

	q = "region=*&order=desc"
	entries, _, status = get[core.Output, Entry](nil, nil, uri.BuildValues(q), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q, status, entries)

	entries, _, status = get[core.Output, Entry](nil, nil, uri.BuildValues(q1), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q1, status, len(entries))

	entries, _, status = get[core.Output, Entry](nil, nil, uri.BuildValues(q2), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q2, status, len(entries))

	entries, _, status = get[core.Output, Entry](nil, nil, uri.BuildValues(q3), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q3, status, len(entries))

	entries, _, status = get[core.Output, Entry](nil, nil, uri.BuildValues(q4), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q4, status, len(entries))

	//Output:
	//test: Get("region=*") -> [status:OK] [entries:4]
	//test: Get("region=*&order=desc") -> [status:OK] [entries:[{4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com case-officer:007  0001-01-01 00:00:00 +0000 UTC} {3 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az3  www.host1.com case-officer:007  0001-01-01 00:00:00 +0000 UTC} {2 director-1 2024-06-10 09:00:35 +0000 UTC us-west-1 usw1-az2  www.host2.com case-officer:006  0001-01-01 00:00:00 +0000 UTC} {1 director-1 2024-06-10 09:00:35 +0000 UTC us-west-1 usw1-az1  www.host1.com case-officer:006  0001-01-01 00:00:00 +0000 UTC}]]
	//test: Get("region=us-west-1&zone=usw1-az1&host=www.host1.com") -> [status:OK] [entries:1]
	//test: Get("region=us-west-1&zone=usw1-az2&host=www.host2.com") -> [status:OK] [entries:1]
	//test: Get("region=us-west-2&zone=usw2-az3&host=www.host1.com") -> [status:OK] [entries:1]
	//test: Get("region=us-west-2&zone=usw2-az4&host=www.host2.com") -> [status:OK] [entries:1]

}

func ExampleGet_Detail() {
	entries, _, status := get[core.Output, EntryDetail](nil, nil, uri.BuildValues(q1), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q1, status, len(entries))

	entries, _, status = get[core.Output, EntryDetail](nil, nil, uri.BuildValues(q2), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q2, status, len(entries))

	entries, _, status = get[core.Output, EntryDetail](nil, nil, uri.BuildValues(q3), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q3, status, len(entries))

	entries, _, status = get[core.Output, EntryDetail](nil, nil, uri.BuildValues(q4), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q4, status, len(entries))

	//Output:
	//test: Get("region=us-west-1&zone=usw1-az1&host=www.host1.com") -> [status:OK] [entries:2]
	//test: Get("region=us-west-1&zone=usw1-az2&host=www.host2.com") -> [status:OK] [entries:1]
	//test: Get("region=us-west-2&zone=usw2-az3&host=www.host1.com") -> [status:Not Found] [entries:0]
	//test: Get("region=us-west-2&zone=usw2-az4&host=www.host2.com") -> [status:Not Found] [entries:0]

}

func ExampleGet_Status() {
	entries, _, status := get[core.Output, EntryStatus](nil, nil, uri.BuildValues(q1), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q1, status, len(entries))

	entries, _, status = get[core.Output, EntryStatus](nil, nil, uri.BuildValues(q2), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q2, status, len(entries))

	entries, _, status = get[core.Output, EntryStatus](nil, nil, uri.BuildValues(q3), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q3, status, len(entries))

	entries, _, status = get[core.Output, EntryStatus](nil, nil, uri.BuildValues(q4), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q4, status, len(entries))

	//Output:
	//test: Get("region=us-west-1&zone=usw1-az1&host=www.host1.com") -> [status:OK] [entries:2]
	//test: Get("region=us-west-1&zone=usw1-az2&host=www.host2.com") -> [status:Not Found] [entries:0]
	//test: Get("region=us-west-2&zone=usw2-az3&host=www.host1.com") -> [status:OK] [entries:1]
	//test: Get("region=us-west-2&zone=usw2-az4&host=www.host2.com") -> [status:OK] [entries:1]

}

func ExampleGet_Update() {
	entries, _, status := get[core.Output, EntryStatusUpdate](nil, nil, uri.BuildValues(q1), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q1, status, len(entries))

	entries, _, status = get[core.Output, EntryStatusUpdate](nil, nil, uri.BuildValues(q2), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q2, status, len(entries))

	entries, _, status = get[core.Output, EntryStatusUpdate](nil, nil, uri.BuildValues(q3), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q3, status, len(entries))

	entries, _, status = get[core.Output, EntryStatusUpdate](nil, nil, uri.BuildValues(q4), "", "", nil)
	fmt.Printf("test: Get(\"%v\") -> [status:%v] [entries:%v]\n", q4, status, len(entries))

	//Output:
	//test: Get("region=us-west-1&zone=usw1-az1&host=www.host1.com") -> [status:OK] [entries:2]
	//test: Get("region=us-west-1&zone=usw1-az2&host=www.host2.com") -> [status:Not Found] [entries:0]
	//test: Get("region=us-west-2&zone=usw2-az3&host=www.host1.com") -> [status:Not Found] [entries:0]
	//test: Get("region=us-west-2&zone=usw2-az4&host=www.host2.com") -> [status:Not Found] [entries:0]

}
