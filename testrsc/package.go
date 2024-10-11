package testrsc

const (
	PkgPath = "github/advanced-go/activity/testrsc"

	COMPath = "file:///f:/files/common"

	Http404Resp = COMPath + "/http-404-resp.txt"
	Http500Resp = COMPath + "/http-500-resp.txt"
	Http504Resp = COMPath + "/http-504-resp.txt"

	CUSTV1Path = "file:///f:/files/customer1"

	CustomerV1Entry             = CUSTV1Path + "/entry.json"
	CustomerGetV1EgressD001Req  = CUSTV1Path + "/get-v1-egress-D001-req.txt"
	CustomerGetV1EgressD001Resp = CUSTV1Path + "/get-v1-egress-D001-resp.txt"

	CustomerGetV1IngressD002Req  = CUSTV1Path + "/get-v1-ingress-D002-req.txt"
	CustomerGetV1IngressD002Resp = CUSTV1Path + "/get-v1-ingress-D002-resp.txt"

	AddrGetV1C001Resp = "file:///f:/files/upstream/customer/get-C001-req-resp.txt"
	AddrGetV1D001Resp = "file:///f:/files/upstream/customer/get-D001-req-resp.txt"

	EventsGetV1LogEgressD001Resp = "file:///f:/files/upstream/events/get-v1-log-egress-D001-req-resp.txt"
)
