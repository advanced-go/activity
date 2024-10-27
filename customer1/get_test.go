package customer1

import (
	"fmt"
	"github.com/advanced-go/activity/testrsc"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/uri"
	"net/url"
)

func ExampleGet_Customer_Failure() {
	values := make(url.Values)
	values.Add(customerKey, "D001")
	path := uri.BuildPath(CustomerAuthority, CustomerV1AddressPath, values)
	h := uri.AddResolverEntry(nil, path, testrsc.Http500Resp)

	path = uri.BuildPath(LogAuthority, LogV1EventPath, values)
	uri.AddResolverEntry(h, path, testrsc.EventsGetV1LogEgressD001Resp)

	h.Add(core.XRequestId, "123-456")
	entries, _, status := get[core.Output](nil, h, logPath, values)
	fmt.Printf("test: get() -> [status:%v] %v\n", status, entries)

	//Output:
	//test: get() -> [status:Internal Error [multiple errors]] []

}

func ExampleGet_Events_Failure() {
	values := make(url.Values)
	values.Add(customerKey, "D001")
	path := uri.BuildPath(CustomerAuthority, CustomerV1AddressPath, values)
	h := uri.AddResolverEntry(nil, path, testrsc.AddrGetV1D001Resp)

	path = uri.BuildPath(LogAuthority, LogV1EventPath, values)
	uri.AddResolverEntry(h, path, testrsc.Http504Resp)

	h.Add(core.XRequestId, "123-456")
	entries, _, status := get[core.Output](nil, h, logPath, values)
	fmt.Printf("test: get() -> [status:%v] %v\n", status, entries)

	//Output:
	//test: get() -> [status:Internal Error [multiple errors]] []

}

func ExampleGet_Customer_Not_Found() {
	values := make(url.Values)
	values.Add(customerKey, "D001")
	path := uri.BuildPath(CustomerAuthority, CustomerV1AddressPath, values)
	h := uri.AddResolverEntry(nil, path, testrsc.Http404Resp)

	path = uri.BuildPath(LogAuthority, LogV1EventPath, values)
	uri.AddResolverEntry(h, path, testrsc.EventsGetV1LogEgressD001Resp)

	h.Add(core.XRequestId, "123-456")
	entries, _, status := get[core.Output](nil, h, logPath, values)
	fmt.Printf("test: get() -> [status:%v] %v\n", status, entries)

	//Output:
	//test: get() -> [status:Not Found] []

}

func ExampleGet_Events_Not_Found() {
	values := make(url.Values)
	values.Add(customerKey, "D001")
	path := uri.BuildPath(CustomerAuthority, CustomerV1AddressPath, values)
	h := uri.AddResolverEntry(nil, path, testrsc.AddrGetV1D001Resp)

	path = uri.BuildPath(LogAuthority, LogV1EventPath, values)
	uri.AddResolverEntry(h, path, testrsc.Http404Resp)

	h.Add(core.XRequestId, "123-456")
	entries, _, status := get[core.Output](nil, h, logPath, values)
	fmt.Printf("test: get() -> [status:%v] %v\n", status, entries)

	//Output:
	//test: get() -> [status:OK] [{{D001 123 Main  Anytown OH 12345 before-email@hotmail.com} []}]

}

func ExampleGet_OK() {
	values := make(url.Values)
	values.Add(customerKey, "D001")
	path := uri.BuildPath(CustomerAuthority, CustomerV1AddressPath, values)
	h := uri.AddResolverEntry(nil, path, testrsc.AddrGetV1D001Resp)

	path = uri.BuildPath(LogAuthority, LogV1EventPath, values)
	uri.AddResolverEntry(h, path, testrsc.EventsGetV1LogEgressD001Resp)

	h.Add(core.XRequestId, "123-456")
	entries, _, status := get[core.Output](nil, h, logPath, values)
	fmt.Printf("test: get() -> [status:%v] %v\n", status, entries)

	//Output:
	//test: get() -> [status:OK] [{{D001 123 Main  Anytown OH 12345 before-email@hotmail.com} [{{us-west oregon dc1 www.test-host.com google-search 123456} 2024-06-03 18:29:16.0447249 +0000 UTC 100 egress GET  200 500 100 10 RL}]}]

}
