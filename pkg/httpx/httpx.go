package httpx

import (
	"github.com/wylswz/cue-se/internal/pkg"
	"io"
	"net/http"
	"net/url"
)

var (
	c = http.Client{}
)

type HttpResponse struct {
	Code    int                 `json:"code"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

func jsonReqHeader(h map[string][]string) map[string][]string {
	if h == nil {
		h = map[string][]string{}
	}
	h["content-type"] = []string{"application/json"}
	return h
}

func JsonGet(urlStr string, parameters pkg.Struct, headers pkg.Struct) (*HttpResponse, error) {
	return doRequest("GET", urlStr, structToMultiValMap(parameters), nil, jsonReqHeader(structToMultiValMap(headers)))
}

func doRequest(method string, urlStr string, parameters map[string][]string, body interface{}, headers map[string][]string) (*HttpResponse, error) {
	urlObj, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	q := urlObj.Query()
	for k, vs := range parameters {
		for _, v := range vs {
			q.Add(k, v)
		}
	}
	urlObj.RawQuery = q.Encode()
	resp, err := c.Do(&http.Request{
		URL:    urlObj,
		Method: method,
		Header: headers,
	})
	if err != nil {
		return nil, err
	}

	bodyBts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		Code:    resp.StatusCode,
		Body:    string(bodyBts),
		Headers: map[string][]string(resp.Header),
	}, nil

}

func structToMultiValMap(pkg.Struct) map[string][]string {
	return nil
}
