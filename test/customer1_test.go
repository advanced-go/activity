package test

import (
	"github.com/advanced-go/activity/customer1"
	http2 "github.com/advanced-go/activity/http"
	"github.com/advanced-go/activity/testrsc"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/core/coretest"
	httpt "github.com/advanced-go/stdlib/httpx/httpxtest"
	"net/http"
	"reflect"
	"testing"
)

func TestCustomer1(t *testing.T) {
	tests := []struct {
		name   string
		req    *http.Request
		resp   *http.Response
		status *core.Status
	}{
		{name: "get-entry", req: httpt.NewRequestTest(testrsc.Customer1EgressGetReq, t), resp: httpt.NewResponseTest(testrsc.Customer1EgressGetResp, t), status: core.StatusOK()},

		//
	}
	for _, tt := range tests {
		ok := true
		t.Run(tt.name, func(t *testing.T) {
			resp, status := http2.Exchange(tt.req)
			if tt.status != nil && status.Code != tt.status.Code {
				t.Errorf("Exchange() got status : %v, want status : %v, error : %v", status.Code, tt.status.Code, status.Err)
				ok = false
			}
			if ok && resp.StatusCode != tt.resp.StatusCode {
				t.Errorf("Exchange() got status code : %v, want status code : %v", resp.StatusCode, tt.resp.StatusCode)
				ok = false
			}
			var gotT []customer1.Entry
			var wantT []customer1.Entry
			if ok {
				gotT, wantT, ok = httpt.Deserialize[coretest.Output, []customer1.Entry](resp.Body, tt.resp.Body, t)
			}
			if ok {
				if !reflect.DeepEqual(gotT, wantT) {
					t.Errorf("Exchange() got = %v, want %v", gotT, wantT)
				}
			}
		})
	}
}
