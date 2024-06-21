package assignment

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"time"
)

func ExampleNullDate() {
	ts := time.Time{}
	year, month, day := ts.Date()
	fmt.Printf("test: timeDate(null) -> [year:%v] [month:%v] [day:%v]\n", year, month, day)

	ts = time.Now().UTC()
	year, month, day = ts.Date()
	fmt.Printf("test: timeDate(Now().UTC()) -> [year:%v] [month:%v] [day:%v]\n", year, month, day)

	//Output:
	//test: timeDate(null) -> [year:1] [month:January] [day:1]
	//test: timeDate(Now().UTC()) -> [year:2024] [month:June] [day:20]

}

func _ExampleAddStatus() {
	o := core.Origin{Region: "us-east-1", Zone: "use1-az1", Host: "www.test.host"}
	st := ""
	agentId := ""
	assigneeId := ""

	status := addStatus(o, st, agentId, assigneeId)
	fmt.Printf("test: addStatus() -> [status:%v]\n", status)

	agentId = "test-agent-id"
	assigneeId = "test-assignee-id"
	o = core.Origin{Region: "us-west-2", Zone: "usw2-az4", Host: "www.host2.com"}
	status = addStatus(o, st, agentId, assigneeId)
	fmt.Printf("test: addStatus() -> [status:%v]\n", status)

	agentId = "test-agent-id"
	assigneeId = "test-assignee-id"
	st = ClosedStatus
	o = core.Origin{Region: "us-west-2", Zone: "usw2-az4", Host: "www.host2.com"}
	status = addStatus(o, st, agentId, assigneeId)
	fmt.Printf("test: addStatus() -> [status:%v] [item:%v]\n", status, lastStatus())

	//Output:
	//test: addStatus() -> [status:Not Found]
	//test: addStatus() -> [status:Bad Request]
	//test: addStatus() -> [status:OK] [item:{4 5 test-agent-id 2024-06-20 17:57:06.1565409 +0000 UTC closed test-assignee-id}]

}

/*
	func ExampleLastStatusFilter() {
		item, ok := lastStatusFilter(-1, "")
		fmt.Printf("test: lastStatusFilter() -> [item:%v] [ok:%v]\n", item, ok)

		item, ok = lastStatusFilter(1, "")
		fmt.Printf("test: lastStatusFilter() -> [item:%v] [ok:%v]\n", item, ok)

		item, ok = lastStatusFilter(3, OpenStatus)
		fmt.Printf("test: lastStatusFilter() -> [item:%v] [ok:%v]\n", item, ok)

		item, ok = lastStatusFilter(3, ClosingStatus)
		fmt.Printf("test: lastStatusFilter() -> [item:%v] [ok:%v]\n", item, ok)

		//Output:
		//test: lastStatusFilter() -> [item:{0 0  0001-01-01 00:00:00 +0000 UTC  }] [ok:false]
		//test: lastStatusFilter() -> [item:{0 0  0001-01-01 00:00:00 +0000 UTC  }] [ok:false]
		//test: lastStatusFilter() -> [item:{0 0  0001-01-01 00:00:00 +0000 UTC  }] [ok:false]
		//test: lastStatusFilter() -> [item:{3 3 agent-name:agent-class:instance-id 2024-06-10 09:00:35 +0000 UTC closing }] [ok:true]

}
*/
func _ExampleAssign() {
	o := core.Origin{Region: "us-west-2", Zone: "usw2-az4", Host: "www.host2.com"}

	status := assign(o, "agent-id", "assignee-id")
	fmt.Printf("test: assign() -> [status:%v] [entry:%v] [item:%v]\n", status, lastEntry(), lastStatus())

	//Output:
	//test: assign() -> [status:OK] [entry:{4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com us-west-2:usw2-az4:case-officer-007  0001-01-01 00:00:00 +0000 UTC assigned}] [item:{4 5 agent-id 2024-06-20 18:11:07.947146 +0000 UTC assigned assignee-id}]

}

func _ExampleAddClosingStatusChange() {
	o := core.Origin{Region: "us-west-2", Zone: "usw2-az4", Host: "www.host2.com"}

	status := addClosingStatusChange(o, "agent-id", "assignee-tag")
	fmt.Printf("test: addClosingStatusChange() -> [status:%v] [entry:%v] [change:%v] [status:%v]\n", status, lastEntry(), lastChange(), lastStatus())

	chg, status1 := getStatusChange(ClosingStatus, "assignee-tag")
	fmt.Printf("test: getStatusChange() -> [status:%v] [change:%v]\n", status1, chg)

	//Output:
	//test: addClosingStatusChange() -> [status:OK] [entry:{4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com us-west-2:usw2-az4:case-officer-007  0001-01-01 00:00:00 +0000 UTC closing}] [change:{4 3 agent-id 2024-06-20 18:30:04.9832473 +0000 UTC assignee-tag closing   0001-01-01 00:00:00 +0000 UTC}] [status:{4 5 agent-id 2024-06-20 18:30:04.9832473 +0000 UTC closing assignee-tag}]
	//test: getStatusChange() -> [status:OK] [item:{4 3 agent-id 2024-06-20 18:25:54.2077016 +0000 UTC assignee-tag closing   0001-01-01 00:00:00 +0000 UTC}]

}

func _ExampleAddReassigningStatusChange() {
	o := core.Origin{Region: "us-west-2", Zone: "usw2-az4", Host: "www.host2.com"}

	status := addReassigningStatusChange(o, "agent-id", "assignee-tag", "new-assignee-tag")
	fmt.Printf("test: addReassigningStatusChange() -> [status:%v] [entry:%v] [change:%v] [status:%v]\n", status, lastEntry(), lastChange(), lastStatus())

	chg, status1 := getStatusChange(ReassigningStatus, "assignee-tag")
	fmt.Printf("test: getStatusChange() -> [status:%v] [change:%v]\n", status1, chg)

	//Output:
	//test: addClosingStatusChange() -> [status:OK] [entry:{4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com us-west-2:usw2-az4:case-officer-007  0001-01-01 00:00:00 +0000 UTC closing}] [change:{4 3 agent-id 2024-06-20 18:30:04.9832473 +0000 UTC assignee-tag closing   0001-01-01 00:00:00 +0000 UTC}] [status:{4 5 agent-id 2024-06-20 18:30:04.9832473 +0000 UTC closing assignee-tag}]
	//test: getStatusChange() -> [status:OK] [item:{4 3 agent-id 2024-06-20 18:25:54.2077016 +0000 UTC assignee-tag closing   0001-01-01 00:00:00 +0000 UTC}]

}
