package testrsc

import (
	"fmt"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/httpx/httpxtest"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"strings"
)

const (
	eventsDir   = "file://[cwd]/files/upstream/events"
	customerDir = "file://[cwd]/files/upstream/customer"
)

type exchange struct {
	dir, req, resp string
}

func requestPath(ex exchange) string {
	return ex.dir + "/" + ex.req
}

func responsePath(ex exchange) string {
	if ex.resp != "" {
		return ex.dir + "/" + ex.resp
	}
	s := strings.Replace(ex.req, ".", "-resp.", 1)
	return ex.dir + "/" + s
}

func newUrl(req *http.Request) string {
	scheme := "https"
	host := req.Host
	if strings.Contains(host, "localhost") {
		scheme = "http"
	}
	return scheme + "://" + host + req.URL.String()
}

func newRequest(req *http.Request) (*http.Request, error) {
	r, err := http.NewRequest(req.Method, newUrl(req), req.Body)
	if err != nil {
		return nil, err
	}
	r.Header = req.Header
	return r, nil
}

func ExampleBuildUpstream() {
	items := []exchange{
		{dir: customerDir, req: "get-req.txt", resp: ""},
		{dir: customerDir, req: "get-all-req.txt", resp: ""},

		{dir: eventsDir, req: "log-egress-v1-get-all-req.txt", resp: ""},
		{dir: eventsDir, req: "log-egress-v2-get-all-req.txt", resp: ""},
		{dir: eventsDir, req: "log-ingress-v1-get-all-req.txt", resp: ""},
		{dir: eventsDir, req: "log-ingress-v2-get-all-req.txt", resp: ""},

		//
	}

	fmt.Printf("BuildUpstream() - start\n")

	for _, i := range items {
		var err error

		req, status := httpxtest.NewRequest(requestPath(i))
		if !status.OK() {
			fmt.Printf("error: NewRequest(\"%v\") ->  %v\n", i.req, status)
			continue
		}
		req, err = newRequest(req)
		if err != nil {
			fmt.Printf("error: updateUrl(\"%v\") -> %v\n", i.req, err)
			continue
		}
		resp, status1 := httpx.Do(req)
		if !status1.OK() {
			fmt.Printf("error: Do(\"%v\") -> %v\n", i.req, status1)
			continue
		}
		if resp.Header.Get(httpx.ContentType) == httpx.ContentTypeJson {
			resp.Body, status = json2.Indent(resp.Body, resp.Header, "", "  ")
			if !status.OK() {
				fmt.Printf("error: json2.Indent(\"%v\") -> %v\n", i.req, status)
				continue
			}
		}
		status = httpxtest.WriteResponse(responsePath(i), resp)
		if !status.OK() {
			fmt.Printf("error: WriteResponse(\"%v\") -> %v\n", i.req, status)
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
