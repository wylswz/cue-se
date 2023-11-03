package httpx

import (
	"encoding/json"
	"github.com/wylswz/cue-se/cue"
	"github.com/wylswz/cue-se/pkg/httpx/clientprovider"
	"io"
	"net/http"
	"net/url"
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

func JsonGet(urlStr string, parametersV cue.Value, headersV cue.Value) (*HttpResponse, error) {
	params, err := structToMultiValMap(parametersV)
	if err != nil {
		return nil, err
	}
	headers, err := structToMultiValMap(headersV)
	if err != nil {
		return nil, err
	}

	return doRequest("GET", urlStr, params, nil, jsonReqHeader(headers))
}

func doRequest(method string, urlStr string, parameters map[string][]string, body interface{}, headers map[string][]string) (*HttpResponse, error) {
	c := clientprovider.C()
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

func structToMultiValMap(s cue.Value) (map[string][]string, error) {
	bts, err := s.MarshalJSON()
	if err != nil {
		return nil, err
	}
	res := map[string][]string{}
	err = json.Unmarshal(bts, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
