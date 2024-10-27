package common

import (
	"fmt"
	"github.com/advanced-go/common/uri"
)

func ExampleValidEntry() {
	q := "region=us-west-1&zone=usw1-az1&host=www.host1.com"
	values := uri.BuildValues(q)
	valid := ValidEntry(values, entryData[0])
	fmt.Printf("test: ValidEntry(\"%v\") -> [valid:%v]\n", q, valid)

	q = "region=us-west-1&zone=usw1-az4&host=www.host1.com"
	values = uri.BuildValues(q)
	valid = ValidEntry(values, entryData[3])
	fmt.Printf("test: ValidEntry(\"%v\") -> [valid:%v]\n", q, valid)

	//Output:
	//test: ValidEntry("region=us-west-1&zone=usw1-az1&host=www.host1.com") -> [valid:true]
	//test: ValidEntry("region=us-west-1&zone=usw1-az4&host=www.host1.com") -> [valid:false]

}
