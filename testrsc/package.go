package testrsc

const (
	PkgPath = "github/advanced-go/customer/testrsc"

	Addr1BasePath = "file:///f:/files/address1"

	Addr1Entry     = Addr1BasePath + "/entry.json"
	Addr1EntryTest = Addr1BasePath + "/entry-test.json"

	Addr1GetReq      = Addr1BasePath + "/get-req.txt"
	Addr1GetResp     = Addr1BasePath + "/get-resp.txt"
	Addr1GetRespTest = Addr1BasePath + "/get-resp-test.txt"

	Addr1GetReqErrHeader   = Addr1BasePath + "/get-req-error-header.txt"
	Addr1GetRespErrContent = Addr1BasePath + "/get-resp-error-content.txt"

	Addr1PutRespFailure = Addr1BasePath + "/put-resp-failure.txt"
)
