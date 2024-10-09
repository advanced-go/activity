package testrsc

const (
	PkgPath = "github/advanced-go/activity/testrsc"

	CUSTV1BasePath = "file:///f:/files/customer1"

	CustomerV1Entry             = CUSTV1BasePath + "/entry.json"
	CustomerGetV1EgressD001Req  = CUSTV1BasePath + "/get-v1-egress-D001-req.txt"
	CustomerGetV1EgressD001Resp = CUSTV1BasePath + "/get-v1-egress-D001-resp.txt"

	CustomerGetV1IngressD002Req  = CUSTV1BasePath + "/get-v1-ingress-D002-req.txt"
	CustomerGetV1IngressD002Resp = CUSTV1BasePath + "/get-v1-ingress-D002-resp.txt"

	AddrGetV1C001Resp = "file:///f:/files/upstream/customer/get-C001-req-resp.txt"
	AddrGetV1D001Resp = "file:///f:/files/upstream/customer/get-D001-req-resp.txt"

	EventsGetV1LogEgressD001Resp = "file:///f:/files/upstream/events/get-v1-log-egress-D001-req-resp.txt"
)
