package telegraph

import (
	"encoding/json"
	"errors"
	"strings"

	jsoniter "github.com/json-iterator/go"
	http "github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

// Response contains a JSON object, which always has a Boolean field ok. If ok equals true, the request was
// successful, and the result of the query can be found in the result field. In case of an unsuccessful request, ok
// equals false, and the error is explained in the error field (e.g. SHORT_NAME_REQUIRED).
type Response struct {
	Ok     bool            `json:"ok"`
	Error  string          `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

var parser = jsoniter.ConfigFastest //nolint:gochecknoglobals
var c = &http.Client{}

// SetSocksProxy change the dialer in http client
func SetSocksProxy(proxy string) {
	c = &http.Client{
		Dial: fasthttpproxy.FasthttpSocksDialer(proxy),
	}
}

// SetHTTPProxy set dialer in http client to a http proxy
func SetHTTPProxy(proxy string) {
	c = &http.Client{
		Dial: fasthttpproxy.FasthttpHTTPDialer(proxy),
	}
}

func makeRequest(path string, payload interface{}) ([]byte, error) {
	src, err := parser.Marshal(payload)
	if err != nil {
		return nil, err
	}

	u := http.AcquireURI()
	defer http.ReleaseURI(u)
	u.SetScheme("https")
	u.SetHost("api.telegra.ph")
	u.SetPath(path)

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.SetRequestURIBytes(u.FullURI())
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetUserAgent("TechMinerApps/telegraph")
	req.Header.SetContentType("application/json")
	req.SetBody(src)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	if err := c.Do(req, resp); err != nil {
		return nil, err
	}

	r := new(Response)
	if err := parser.Unmarshal(resp.Body(), r); err != nil {
		return nil, err
	}

	if !r.Ok {
		if strings.Contains(r.Error, "FLOOD_WAIT") {
			return nil, ErrFloodWait
		}
		return nil, errors.New(r.Error)
	}

	return r.Result, nil
}
