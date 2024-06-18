package inference

import (
	"fmt"
	"github.com/advanced-go/stdlib/uri"
)

func _ExampleOrder() {
	q := ""
	result := order(nil, entryData)
	fmt.Printf("test: order(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), result)

	q = "order=desc"
	result = order(uri.BuildValues(q), entryData)
	fmt.Printf("test: order(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), result)

	//Output:
	//fail
}

func ExampleTop() {
	q := ""
	result := top(nil, entryData)
	fmt.Printf("test: top(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), len(result))

	q = "top=1"
	result = top(uri.BuildValues(q), entryData)
	fmt.Printf("test: top(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), len(result))

	//Output:
	//test: top("") -> [cnt:2] [result:2]
	//test: top("top=1") -> [cnt:2] [result:1]

}

func ExampleDistinct() {
	q := ""
	result := distinct(nil, entryData)
	fmt.Printf("test: distinct(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), len(result))

	q = "distinct=host"
	result = distinct(uri.BuildValues(q), entryData)
	fmt.Printf("test: distinct(\"%v\") -> [cnt:%v] [result:%v]\n", q, len(entryData), len(result))

	//Output:
	//test: distinct("") -> [cnt:2] [result:2]
	//test: distinct("distinct=host") -> [cnt:2] [result:2]

}
