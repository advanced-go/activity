package test

import (
	http2 "github.com/advanced-go/customer/http"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"reflect"
	"testing"
)

func TestExchange(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name  string
		args  args
		want  *http.Response
		want1 *core.Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := http2.Exchange(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exchange() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Exchange() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
