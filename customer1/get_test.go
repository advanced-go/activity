package customer1

import (
	"fmt"
	"github.com/advanced-go/activity/testrsc"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/url"
)

func ExampleGet_Customer() {
	values := make(url.Values)
	values.Add(customerKey, "D001")
	path := uri.BuildPath(CustomerAuthority, Customer1AddressPath, values)
	h := uri.AddResolverContentLocation(nil, path, testrsc.CustomerD001GetResp)

	path = uri.BuildPath(EventsAuthority, Events1EgressPath, values)
	uri.AddResolverContentLocation(h, path, testrsc.EventsV1LogEgressD001GetResp)

	h.Add(core.XRequestId, "123-456")
	entries, _, status := get[core.Output](nil, h, activity1EgressPath, values)
	fmt.Printf("test: get() -> [status:%v] %v\n", status, entries)

	//Output:
	//test: get() -> [status:OK] [{{D001 123 Main  Anytown OH 12345 before-email@hotmail.com} [{{us-west oregon dc1 www.test-host.com google-search 123456} 2024-06-03 18:29:16.0447249 +0000 UTC 100 egress GET  200 500 100 10 RL}]}]

}

/*
func ExampleGet_Customer_All() {
	values := make(url.Values)
	values.Add(customerKey, "*")
	path := uri.BuildPath("", StoragePath, values)
	h := uri.AddResolverContentLocation(nil, path, testrsc.Addr1GetRespTest)

	entries, _, status := get[core.Output](nil, h, values)
	fmt.Printf("test: get() -> [status:%v] [path:%v] [entries:%v]\n", status, path, len(entries))

	//Output:
	//test: get() -> [status:OK] [path:storage/address?customer=*] [entries:4]

}

func ExampleGet_State() {
	values := make(url.Values)
	values.Add(stateKey, "IA")
	path := uri.BuildPath("", StoragePath, values)
	h := uri.AddResolverContentLocation(nil, path, testrsc.Addr1GetRespTest)

	entries, _, status := get[core.Output](nil, h, values)
	fmt.Printf("test: get() -> [status:%v] [path:%v] [entries:%v]\n", status, path, len(entries))

	//Output:
	//test: get() -> [status:OK] [path:storage/address?state=IA] [entries:2]

}


*/
