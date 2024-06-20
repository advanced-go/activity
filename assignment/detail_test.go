package assignment

import "fmt"

func ExampleLastDetail() {

	fmt.Printf("test: lastDetail() -> [entry:%v] [detail:%v]\n", lastDetail().EntryId, lastDetail().DetailId)

	//Output:
	//test: lastDetail() -> [entry:2] [detail:3]
	
}
