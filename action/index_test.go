package action

import (
	"fmt"
	"github.com/advanced-go/stdlib/uri"
)

func ExampleGet_Index() {
	fmt.Printf("test: init() -> [index:%v]\n", index)

	q := "region=us-west-1&zone=usw1-az1&host=www.host1.com"
	values := uri.BuildValues(q)
	_, exists := lookupEntry(values)
	fmt.Printf("test: entryExists(\"%v\") -> [exists:%v]\n", q, exists)

	q = "region=us-west-1&zone=usw1-az4&host=www.host1.com"
	values = uri.BuildValues(q)
	_, exists = lookupEntry(values)
	fmt.Printf("test: entryExists(\"%v\") -> [exists:%v]\n", q, exists)

	//Output:
	//test: init() -> [index:map[us-west-1:usw1-az1:www.host1.com:{1 director-1 2024-06-10 09:00:35 +0000 UTC us-west-1 usw1-az1  www.host1.com case-officer:006  0001-01-01 00:00:00 +0000 UTC} us-west-1:usw1-az2:www.host2.com:{2 director-1 2024-06-10 09:00:35 +0000 UTC us-west-1 usw1-az2  www.host2.com case-officer:006  0001-01-01 00:00:00 +0000 UTC} us-west-2:usw2-az3:www.host1.com:{3 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az3  www.host1.com case-officer:007  0001-01-01 00:00:00 +0000 UTC} us-west-2:usw2-az4:www.host2.com:{4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com case-officer:007  0001-01-01 00:00:00 +0000 UTC}]]
	//test: entryExists("region=us-west-1&zone=usw1-az1&host=www.host1.com") -> [exists:true]
	//test: entryExists("region=us-west-1&zone=usw1-az4&host=www.host1.com") -> [exists:false]

}
