package telegraph

import (
	"errors"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	http "github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

var parser = jsoniter.ConfigFastest //nolint:gochecknoglobals

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

func (c *client) makeRequest(path string, payload interface{}) ([]byte, error) {
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

	if err := c.httpClient.Do(req, resp); err != nil {
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

type Client interface {
	Client() *http.Client
	ContentFormat(data interface{}) (n []Node, err error)
	CreateAccount(account Account) (*Account, error)
	CreatePage(page Page, returnContent bool) (*Page, error)
	GetViews(path string, date time.Time) (*PageViews, error)
	GetPage(path string, returnContent bool) (*Page, error)
	RevokeAccessToken() (*Account, error)
}

type client struct {
	httpClient *http.Client
	Account    *Account
}

func (c *client) Client() *http.Client {
	return c.httpClient
}

func NewClient() (Client, error) {
	cc := &client{
		httpClient: &http.Client{},
		Account: &Account{
			AccessToken: "",
			AuthURL:     "",
			ShortName:   "",
			AuthorName:  "",
			AuthorURL:   "",
			PageCount:   0,
		},
	}

	return cc, nil
}
