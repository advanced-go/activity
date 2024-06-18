package action

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
)

func ExamplePut() {
	_, status := put[core.Output, Entry](nil, nil, actionResource, "", nil, nil)
	fmt.Printf("test: put(nil,h,nil) -> [status:%v] [count:%v]\n", status, len(entryData))

	_, status = put[core.Output, Entry](nil, nil, actionResource, "", []Entry{{Region: "us-east"}}, nil)
	fmt.Printf("test: put(nil,h,[]Entry) -> [status:%v] [count:%v]\n", status, len(entryData))

	//Output:
	//test: put(nil,h,nil) -> [status:Invalid Content [error: no entries found]] [count:4]
	//test: put(nil,h,[]Entry) -> [status:OK] [count:5]

}
