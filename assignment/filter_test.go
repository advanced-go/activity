package assignment

import (
	"fmt"
	"github.com/advanced-go/stdlib/uri"
)

func ExampleOrder_Entry() {
	q := ""
	result := Order(nil, entryData)
	fmt.Printf("test: Order(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), result)

	q = "order=desc"
	result = Order(uri.BuildValues(q), entryData)
	fmt.Printf("test: Order(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), result)

	//Output:
	//fail
}

func ExampleTop_Entry() {
	q := ""
	result := Top(nil, entryData)
	fmt.Printf("test: Top(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), len(result))

	q = "top=2"
	result = Top(uri.BuildValues(q), entryData)
	fmt.Printf("test: Top(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), len(result))

	//Output:
	//test: top("") -> [cnt:4] [result:4]
	//test: top("top=2") -> [cnt:4] [result:2]

}
