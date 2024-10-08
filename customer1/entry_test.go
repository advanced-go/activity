package customer1

import (
	"fmt"
	"github.com/advanced-go/activity/testrsc"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
	"time"
)

var data = []Entry{
	{Customer: address{
		CustomerId:   "C001",
		AddressLine1: "Test Line 1",
		AddressLine2: "Test Line 2",
		City:         "City",
		State:        "State",
		PostalCode:   "PostalCode",
		Email:        "email@email.com",
	},
		Activity: []log{{Origin: core.Origin{Region: "us-west", Zone: "oregon", SubZone: "dc1", Host: "www.google.com", Route: "google-search"},
			StartTime: time.Now().UTC(), Duration: 123, Traffic: "Ingress", Method: "GET", Uri: "test", StatusCode: 200, Timeout: 2500, RateLimit: -1, RateBurst: -1, ControllerCode: "TO"},
		},
	},
}

func ExampleEncode() {
	buf, status := json.Marshal(data)

	fmt.Printf("test: Encode() -> [status:%v] %v\n", status, string(buf))

	//Output:
	//fail

}

func ExampleDecode() {
	e, status := json.New[[]Entry](testrsc.CustomerV1Entry, nil)
	fmt.Printf("test: Decode()-> [status:%v] %v\n", status, e)

	//Output:
	//fail

}
