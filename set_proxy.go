package telegraph

import (
	http "github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

// SetSocksProxy change the dialer in http client
func (c *client) SetSocksProxy(proxy string) {
	c.httpClient = &http.Client{
		Dial: fasthttpproxy.FasthttpSocksDialer(proxy),
	}
}

// SetHTTPProxy set dialer in http client to a http proxy
func (c *client) SetHTTPProxy(proxy string) {
	c.httpClient = &http.Client{
		Dial: fasthttpproxy.FasthttpHTTPDialer(proxy),
	}
}
