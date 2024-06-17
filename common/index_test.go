package common

import (
	"fmt"
	"github.com/advanced-go/stdlib/uri"
)

func ExampleOriginIndex_AddEntry() {
	q := "region=us-west-1&zone=usw1-az1&host=www.host1.com"
	status := entryList.AddEntry(entryData[0])

	fmt.Printf("test: AddEntry(\"%v\") -> [status:%v]\n", q, status)

	//Output:
	//test: AddEntry("region=us-west-1&zone=usw1-az1&host=www.host1.com") -> [status:Bad Request]

}

func ExampleOriginIndex_LookupEntry() {
	q := "region=us-west-1&zone=usw1-az1&host=www.host1.com"
	values := uri.BuildValues(q)
	_, exists := entryList.LookupEntry(values)
	fmt.Printf("test: LookupEntry(\"%v\") -> [exists:%v]\n", q, exists)

	q = "region=us-west-1&zone=usw1-az4&host=www.host1.com"
	values = uri.BuildValues(q)
	_, exists = entryList.LookupEntry(values)
	fmt.Printf("test: LookupEntry(\"%v\") -> [exists:%v]\n", q, exists)

	//Output:
	//test: LookupEntry("region=us-west-1&zone=usw1-az1&host=www.host1.com") -> [exists:true]
	//test: LookupEntry("region=us-west-1&zone=usw1-az4&host=www.host1.com") -> [exists:false]

}
