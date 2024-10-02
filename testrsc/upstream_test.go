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
	dirPath = "file://[cwd]/files/upstream"
)

func path(name string) string {
	return dirPath + "/" + name
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
	items := []struct {
		req  string
		resp string
	}{
		{req: "customer-get-req.txt", resp: "customer-get-resp.txt"},

		//
	}
	fmt.Printf("BuildUpstream() - start\n")

	for _, i := range items {
		var err error

		req, status := httpxtest.NewRequest(path(i.req))
		if !status.OK() {
			fmt.Printf("error: NewRequest() ->  %v\n", status)
			continue
		}
		req, err = newRequest(req)
		if err != nil {
			fmt.Printf("error: updateUrl() -> %v\n", err)
			continue
		}
		resp, status1 := httpx.Do(req)
		if !status1.OK() {
			fmt.Printf("error: Do() -> %v\n", status)
			continue
		}
		if resp.Header.Get(httpx.ContentType) == httpx.ContentTypeJson {
			resp.Body, status = json2.Indent(resp.Body, resp.Header, "", "  ")
			if !status.OK() {
				fmt.Printf("error: json2.Indent() -> %v\n", status)
				continue
			}
		}
		status = httpxtest.WriteResponse(path(i.resp), resp)
		if !status.OK() {
			fmt.Printf("error: WriteResponse() -> %v\n", status)
			continue
		}
	}

	fmt.Printf("BuildUpstream() - stop")

	//Output:
	//BuildUpstream() - start
	//BuildUpstream() - stop

}
