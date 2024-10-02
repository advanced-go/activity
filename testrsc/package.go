package testrsc

const (
	PkgPath = "github/advanced-go/activity/testrsc"

	cust1BasePath = "file:///f:/files/customer1"

	Customer1Entry         = cust1BasePath + "/entry.json"
	Customer1EgressGetReq  = cust1BasePath + "/egress-get-req.txt"
	Customer1EgressGetResp = cust1BasePath + "/egress-get-resp.txt"

	CustomerServiceGetResp = cust1BasePath + "/customer-get-resp.txt"
	EventsServiceGetResp   = cust1BasePath + "/egress-log-get-resp.txt"
)
