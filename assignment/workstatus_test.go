package assignment

import (
	"fmt"
	"time"
)

func ExampleNullDate() {
	ts := time.Time{}
	year, month, day := ts.Date()
	fmt.Printf("test: timeDate() -> [year:%v] [month:%v] [day:%v]\n", year, month, day)

	ts = time.Now().UTC()
	year, month, day = ts.Date()
	fmt.Printf("test: timeDate() -> [year:%v] [month:%v] [day:%v]\n", year, month, day)

	//Output:
	//fail

}
