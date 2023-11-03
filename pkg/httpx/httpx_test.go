package httpx_test

import (
	"github.com/wylswz/cue-se/pkg/httpx/clientprovider"
	"github.com/wylswz/cue-se/pkg/internal/builtintest"
	"net/http"
	"strconv"
	"testing"
)

var (
	_ http.RoundTripper = &echoRoundTripper{}
)

type echoRoundTripper struct{}

func status(req *http.Request) int {
	v := req.Header["x-cue-se-status"]
	if len(v) == 0 {
		return 200
	}
	res, _ := strconv.Atoi(v[0])
	return res
}

func (e echoRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: status(request),
		Header:     request.Header,
		Body:       request.Body,
	}, nil
}

func TestHttpx(t *testing.T) {
	clientprovider.ConfigRoundTripper(echoRoundTripper{})
	builtintest.Run("httpx", t)
}
