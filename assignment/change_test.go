package assignment

import "fmt"

func ExampleLastChange() {

	fmt.Printf("test: lastChange() -> [entry:%v] [change:%v]\n", lastChange().EntryId, lastChange().ChangeId)

	//Output:
	//test: lastChange() -> [entry:1] [change:2]

}
