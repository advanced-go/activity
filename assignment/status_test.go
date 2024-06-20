package assignment

import "fmt"

func ExampleLastStatus() {

	fmt.Printf("test: lastStatus() -> [entry:%v] [change:%v]\n", lastStatus().EntryId, lastStatus().StatusId)

	//Output:
	//test: lastStatus() -> [entry:4] [change:4]

}
