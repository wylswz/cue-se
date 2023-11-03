package clientprovider

import "net/http"

var (
	c = &http.Client{}
)

func C() *http.Client {
	return c
}

func ConfigRoundTripper(rt http.RoundTripper) {
	c.Transport = rt
}
