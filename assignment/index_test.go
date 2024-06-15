package assignment

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
	//test: init() -> [index:map[us-west-1:usw1-az1:www.host1.com:1 us-west-1:usw1-az2:www.host2.com:2 us-west-2:usw2-az3:www.host1.com:3 us-west-2:usw2-az4:www.host2.com:4]]
	//test: entryExists("region=us-west-1&zone=usw1-az1&host=www.host1.com") -> [exists:true]
	//test: entryExists("region=us-west-1&zone=usw1-az4&host=www.host1.com") -> [exists:false]

}
