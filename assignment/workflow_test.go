package assignment

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
)

func _ExampleInsert() {
	e := lastEntry()
	o := core.Origin{Region: e.Region, Zone: e.Zone, Host: e.Host}

	status := insert(o, "agent-id", "assignee-tag")
	fmt.Printf("test: insert() -> [status:%v]\n", status)

	o = core.Origin{Region: "us-east-1", Zone: "use1-az1", Host: "www.test.host"}
	status = insert(o, "agent-id", "assignee-id")
	fmt.Printf("test: insert() -> [status:%v] [entry:%v] [status:%v]\n", status, lastEntry(), lastStatus())

	//Output:
	//test: insert() -> [status:Invalid Content [error: assignment already exist for: {us-west-2 usw2-az4  www.host2.com }]]
	//test: insert() -> [status:OK] [entry:{5 agent-id 2024-06-21 10:59:50.1725978 +0000 UTC us-east-1 use1-az1  www.test.host assignee-id  0001-01-01 00:00:00 +0000 UTC }] [status:{5 5 agent-id 2024-06-21 10:59:50.1725978 +0000 UTC open }]

}

func ExampleGetOpen() {
	assigneeTag := "us-west-2:usw2-az4:case-officer-007"
	entries, status := getOpen(assigneeTag)
	fmt.Printf("test: getOpen() -> [status:%v] [entry:%v]\n", status, entries)

	//Output:
	//test: getOpen() -> [status:OK] [entry:[{4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com us-west-2:usw2-az4:case-officer-007  0001-01-01 00:00:00 +0000 UTC open}]]

}

func ExampleGetUpdateEntryStatus() {
	e := lastEntry()
	o := core.Origin{Region: e.Region, Zone: e.Zone, Host: e.Host}

	status := updateEntryStatus(o, ClosingStatus)
	fmt.Printf("test: getOpen() -> [status:%v] [entry:%v]\n", status, lastEntry())

	//Output:
	//test: getOpen() -> [status:OK] [entry:{4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com us-west-2:usw2-az4:case-officer-007  0001-01-01 00:00:00 +0000 UTC closing}]

}

func _ExampleAddDetail() {
	e := lastEntry()
	o := core.Origin{Region: e.Region, Zone: e.Zone, Host: e.Host}

	status := addDetail(o, "agent-id", "route-name", "detail description")
	fmt.Printf("test: addDetail() -> [status:%v] [detail:%v]\n", status, lastDetail())

	//Output:
	//test: addDetail() -> [status:OK] [detail:{4 4 agent-id 2024-06-21 11:18:47.3964939 +0000 UTC route-name detail description}]

}

func _ExampleProcessClose() {
	e := lastEntry()
	o := core.Origin{Region: e.Region, Zone: e.Zone, Host: e.Host}

	chg := lastChange()
	status := processClose(o, "agent-id", chg)
	fmt.Printf("test: processClose() -> [status:%v] [entry:%v] [status:%v] [change:%v]\n", status, lastEntry(), lastStatus(), lastChange())

	//Output:
	//test: processClose() -> [status:OK] [entry:{4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com us-west-2:usw2-az4:case-officer-007  0001-01-01 00:00:00 +0000 UTC closed}] [status:{4 5 agent-id 2024-06-21 11:36:08.2576177 +0000 UTC closed }] [change:{1 2 agent-name:agent-class:instance-id 2024-06-10 09:00:35 +0000 UTC class2 closed new test2 error 2024-06-21 11:36:08.2576177 +0000 UTC}]

}

func _ExampleProcessReassignment() {
	e := lastEntry()
	o := core.Origin{Region: e.Region, Zone: e.Zone, Host: e.Host}

	chg := lastChange()
	status := processReassignment(o, "agent-id", chg)
	fmt.Printf("test: processClose() -> [status:%v] [entry:%v] [status:%v] [change:%v]\n", status, lastEntry(), lastStatus(), lastChange())

	//Output:
	//test: processClose() -> [status:OK] [entry:{4 director-2 2024-06-10 09:00:35 +0000 UTC us-west-2 usw2-az4  www.host2.com us-west-2:usw2-az4:case-officer-007  0001-01-01 00:00:00 +0000 UTC closed}] [status:{4 5 agent-id 2024-06-21 11:36:08.2576177 +0000 UTC closed }] [change:{1 2 agent-name:agent-class:instance-id 2024-06-10 09:00:35 +0000 UTC class2 closed new test2 error 2024-06-21 11:36:08.2576177 +0000 UTC}]

}
