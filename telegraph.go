package telegraph

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	http "github.com/valyala/fasthttp"
)

var parser = jsoniter.ConfigFastest

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

// Client is the interface to act with telegraph API
// The 9 method is identical to the 9 method provided at https://telegra.ph/api
type Client interface {

	// Account returns the account object client is using
	Account() *Account

	// Client return the http client using, not recommend using
	Client() *http.Client

	// ContentFormat format the input of string or byte to Node object
	ContentFormat(data interface{}) (n []Node, err error)

	// CreateAccount is used to create a new account using provided info
	// It alse set the account info stored in client
	CreateAccount(account Account) (*Account, error)

	// CreatePage use the account info stored within client to create a new page
	CreatePage(page Page, returnContent bool) (*Page, error)

	// EditAccountInfo update the account info according to input
	// It ignores the Token field in the input account object and use the one stored within client
	EditAccountInfo(update Account) (*Account, error)

	// EditPage update the Page provided
	EditPage(update Page, returnContent bool) (*Page, error)

	// GetAccountInfo get the account info from the token stored in client object
	// It does not update the account info in client object
	GetAccountInfo(fields ...string) (*Account, error)

	// GetPage get the page object from provided url
	GetPage(path string, returnContent bool) (*Page, error)

	// GetPageList
	GetPageList(offset, limit int) (*PageList, error)

	// GetViews get the views count of specific page
	GetViews(path string, date time.Time) (*PageViews, error)

	// RevokeAccessToken is used when accesstoken is compromises
	RevokeAccessToken() (*Account, error)

	// SetSocksProxy
	SetSocksProxy(proxy string)

	// SetHTTPProxy
	SetHTTPProxy(proxy string)
}

type client struct {
	httpClient *http.Client
	account    *Account
}

func (c *client) Client() *http.Client {
	return c.httpClient
}

func (c *client) Account() *Account {
	return c.account
}

// NewClient return a empty client
// useful when creating new account or testing
// otherwise should use NewClientWithToken
func NewClient() Client {
	c := &client{
		httpClient: &http.Client{},
		account: &Account{
			AccessToken: "",
			AuthURL:     "",
			ShortName:   "",
			AuthorName:  "",
			AuthorURL:   "",
			PageCount:   0,
		},
	}

	if c == nil {
		panic("unexpected newclient object allocation error")
	}
	return c
}

// NewClientWithToken create a client from a valid token
// the client is already usable with account info set correctly
func NewClientWithToken(token string, httpClientList ...*http.Client) (Client, error) {

	// Check parameter length
	length := len(httpClientList)

	// Not allow more than one http client
	if length > 1 {
		return nil, errors.New("too many http client provided")
	}

	// Set the client to the provided one
	if length == 0 {
		c := &client{
			httpClient: &http.Client{},
			account: &Account{
				AccessToken: token,
				AuthURL:     "",
				ShortName:   "",
				AuthorName:  "",
				AuthorURL:   "",
				PageCount:   0,
			},
		}
		var err error
		c.account, err = c.GetAccountInfo()
		if err != nil {
			return nil, fmt.Errorf("error getting account info: %v", err)
		}

		return c, nil
	}

	// Default behaviour
	c := &client{
		httpClient: httpClientList[0],
		account: &Account{
			AccessToken: token,
			AuthURL:     "",
			ShortName:   "",
			AuthorName:  "",
			AuthorURL:   "",
			PageCount:   0,
		},
	}
	var err error
	c.account, err = c.GetAccountInfo()
	if err != nil {
		return nil, fmt.Errorf("error getting account info: %v", err)
	}

	return c, nil
}
