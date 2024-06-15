package http

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"io"
	url2 "net/url"
)

const (
	q1 = "region=us-west-1&zone=usw1-az1&host=www.host1.com"
	q2 = "region=us-west-1&zone=usw1-az2&host=www.host2.com"
	q3 = "region=us-west-2&zone=usw2-az3&host=www.host1.com"
	q4 = "region=us-west-2&zone=usw2-az4&host=www.host2.com"
)

func ExampleAssignmentGet_Entry() {
	url, _ := url2.Parse("http://localhost:8080/github/advanced-go/activity:assignment/entry?" + q1)
	p := uri.Uproot(url.String())

	var buf []byte

	resp, status := assignmentGet[core.Output](nil, nil, url, p)
	if status.OK() {
		buf, _ = io.ReadAll(resp.Body)
		//s = string(buf)
	}
	fmt.Printf("test: assignmentGet() -> [status:%v] [status-code:%v] [body:%v]\n", status, resp.StatusCode, string(buf))

	//Output:
	//test: assignmentGet() -> [status:OK] [status-code:200] [body:[{"entry-id":"1","agent-id":"director-1","created-ts":"2024-06-10T09:00:35Z","region":"us-west-1","zone":"usw1-az1","sub-zone":"","host":"www.host1.com","assignee-class":"case-officer:006","assignee-id":"","updated-ts":"0001-01-01T00:00:00Z"}]]

}

func ExampleAssignmentGet_Detail() {
	url, _ := url2.Parse("http://localhost:8080/github/advanced-go/activity:assignment/detail?" + q1)
	p := uri.Uproot(url.String())

	var buf []byte

	resp, status := assignmentGet[core.Output](nil, nil, url, p)
	if status.OK() {
		buf, _ = io.ReadAll(resp.Body)
		//s = string(buf)
	}
	fmt.Printf("test: assignmentGet() -> [status:%v] [status-code:%v] [body:%v]\n", status, resp.StatusCode, string(buf))

	//Output:
	//test: assignmentGet() -> [status:OK] [status-code:200] [body:[{"entry-id":"1","detail-id":"1","agent-id":"agent-name:agent-class:instance-id","created-ts":"2024-06-10T09:00:35Z","route":"search","details":"various information"},{"entry-id":"1","detail-id":"2","agent-id":"agent-name:agent-class:instance-id","created-ts":"2024-06-10T09:00:35Z","route":"host","details":"other information"}]]

}

func ExampleAssignmentGet_Status() {
	url, _ := url2.Parse("http://localhost:8080/github/advanced-go/activity:assignment/status?" + q1)
	p := uri.Uproot(url.String())

	var buf []byte

	resp, status := assignmentGet[core.Output](nil, nil, url, p)
	if status.OK() {
		buf, _ = io.ReadAll(resp.Body)
		//s = string(buf)
	}
	fmt.Printf("test: assignmentGet() -> [status:%v] [status-code:%v] [body:%v]\n", status, resp.StatusCode, string(buf))

	//Output:
	//test: assignmentGet() -> [status:OK] [status-code:200] [body:[{"entry-id":"1","status-id":"1","agent-id":"agent-name:agent-class:instance-id","created-ts":"2024-06-10T09:00:35Z","status":"open","assignee-id":""},{"entry-id":"1","status-id":"2","agent-id":"agent-name:agent-class:instance-id","created-ts":"2024-06-10T09:00:35Z","status":"closed","assignee-id":""}]]

}

func ExampleAssignmentGet_StatusUpdate() {
	var buf []byte

	url, _ := url2.Parse("http://localhost:8080/github/advanced-go/activity:assignment/update?" + q1)
	p := uri.Uproot(url.String())
	// Bad request, no resource with assignment/update
	resp, status := assignmentGet[core.Output](nil, nil, url, p)
	if status.OK() {
		buf, _ = io.ReadAll(resp.Body)
	}
	fmt.Printf("test: assignmentGet() -> [status:%v] [status-code:%v] [body:%v]\n", status, resp.StatusCode, string(buf))

	buf = nil
	url, _ = url2.Parse("http://localhost:8080/github/advanced-go/activity:assignment/status-update?" + q1)
	p = uri.Uproot(url.String())
	resp, status = assignmentGet[core.Output](nil, nil, url, p)
	if status.OK() {
		buf, _ = io.ReadAll(resp.Body)
	}
	fmt.Printf("test: assignmentGet() -> [status:%v] [status-code:%v] [body:%v]\n", status, resp.StatusCode, string(buf))

	buf = nil
	url, _ = url2.Parse("http://localhost:8080/github/advanced-go/activity:assignment/status-update?" + q2)
	p = uri.Uproot(url.String())
	resp, status = assignmentGet[core.Output](nil, nil, url, p)
	if status.OK() {
		buf, _ = io.ReadAll(resp.Body)
	}
	fmt.Printf("test: assignmentGet() -> [status:%v] [status-code:%v] [body:%v]\n", status, resp.StatusCode, string(buf))

	//Output:
	//test: assignmentGet() -> [status:Bad Request] [status-code:400] [body:]
	//test: assignmentGet() -> [status:OK] [status-code:200] [body:[{"entry-id":"1","update-id":"1","agent-id":"agent-name:agent-class:instance-id","created-ts":"2024-06-10T09:00:35Z","assignee-class":"class","NewStatus":"closed","new-assignee-class":"new","error":"test error","processed-ts":"0001-01-01T00:00:00Z"},{"entry-id":"1","update-id":"2","agent-id":"agent-name:agent-class:instance-id","created-ts":"2024-06-10T09:00:35Z","assignee-class":"class2","NewStatus":"closed","new-assignee-class":"new","error":"test2 error","processed-ts":"0001-01-01T00:00:00Z"}]]
	//test: assignmentGet() -> [status:Not Found] [status-code:404] [body:]

}
