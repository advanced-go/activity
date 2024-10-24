package testrsc

import (
	"fmt"
	"github.com/advanced-go/common/httpx"
	"github.com/advanced-go/common/test"
	"github.com/advanced-go/stdlib/httpx/httpxtest"
)

const (
	eventsDir   = "file://[cwd]/files/upstream/events"
	customerDir = "file://[cwd]/files/upstream/customer"
)

func ExampleBuildUpstream() {
	items := []test.FileList{
		{Dir: customerDir, Req: "get-D002-req.txt"},
		{Dir: customerDir, Req: "get-D001-req.txt"},

		{Dir: eventsDir, Req: "get-v1-log-egress-D001-req.txt"},
		{Dir: eventsDir, Req: "get-v1-log-ingress-D002-req.txt"},

		//
	}

	fmt.Printf("BuildUpstream() - start\n")

	for _, info := range items {
		var err error

		req, status := httpxtest.NewRequest(info.RequestPath())
		if !status.OK() {
			fmt.Printf("error: NewRequest(\"%v\") ->  %v\n", info.Req, status)
			continue
		}
		req, status = info.NewRequest(req)
		if !status.OK() {
			fmt.Printf("error: NewRequest(\"%v\") -> %v\n", info.Req, err)
			continue
		}
		resp, status1 := httpx.Do(req)
		if !status1.OK() {
			fmt.Printf("error: Do(\"%v\") -> [status:%v]] [url:%v]\n", info.Req, status1, req.URL.String())
			continue
		}
		if resp.Header.Get(httpx.ContentType) == httpx.ContentTypeJson {
			resp.Body, status = json2.Indent(resp.Body, resp.Header, "", "  ")
			if !status.OK() {
				fmt.Printf("error: json2.Indent(\"%v\") -> %v\n", info.Req, status)
				continue
			}
		}
		status = httpxtest.WriteResponse(info.ResponsePath(), resp)
		if !status.OK() {
			fmt.Printf("error: WriteResponse(\"%v\") -> %v\n", info.Req, status)
			continue
		}
	}

	fmt.Printf("BuildUpstream() - stop")

	//Output:
	//BuildUpstream() - start
	//BuildUpstream() - stop

}

/*
func ExampleBuildUpstream() {
	items := []exchange {{
		req  string
		resp string
	},
	}{


*/
