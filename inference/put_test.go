package inference

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
)

func ExamplePut() {
	_, status := put[core.Output](nil, nil, nil)
	fmt.Printf("test: put(nil,h,nil) -> [status:%v] [count:%v]\n", status, len(entryData))

	_, status = put[core.Output](nil, nil, []Entry{{Region: "us-east"}})
	fmt.Printf("test: put(nil,h,[]Entry) -> [status:%v] [count:%v]\n", status, len(entryData))

	//Output:
	//test: put(nil,h,nil) -> [status:OK] [count:2]
	//test: put(nil,h,[]Entry) -> [status:OK] [count:3]

}
