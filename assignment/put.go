package assignment

import (
	"bytes"
	"context"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"io"
	"net/http"
)

func put[E core.ErrorHandler](ctx context.Context, h http.Header, body []Entry) (http.Header, *core.Status) {
	//var e E

	if len(body) == 0 {
		return nil, core.StatusOK()
	}
	if ctx == nil {
		ctx = context.Background()
	}
	/*
		rc, _, status := createReadCloser(body)
		if !status.OK() {
			e.Handle(status, core.RequestId(h))
			return nil, status
		}

			url := resolver.Url(hostKey, module.DocumentsAuthority, module.DocumentsResourceV1, nil, h)
			req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, rc)
			if err != nil {
				return nil, core.NewStatusError(core.StatusInvalidArgument, err)
			}
			req.Header.Set(core.XFrom, module.Authority)
			httpx.Forward(req.Header, h)
			_, status = httpx.Exchange(req)
			if !status.OK() {
				e.Handle(status, core.RequestId(h))
			}
			return nil, status

	*/
	storage = append(storage, body...)
	return nil, core.StatusOK()
}

func createReadCloser(body any) (io.ReadCloser, int64, *core.Status) {
	switch ptr := body.(type) {
	case []Entry:
		return json2.NewReadCloser(body)
	case []byte:
		return io.NopCloser(bytes.NewReader(ptr)), int64(len(ptr)), core.StatusOK()
	default:
		return nil, 0, core.NewStatusError(core.StatusInvalidArgument, core.NewInvalidBodyTypeError(body))
	}
}
