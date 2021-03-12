package telegraph

type getAccountInfo struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`

	// List of account fields to return.
	Fields []string `json:"fields,omitempty"`
}

// GetAccountInfo get information about a Telegraph account. Returns an Account object on success.
func (c *client) GetAccountInfo(fields ...string) (*Account, error) {
	data, err := c.makeRequest("getAccountInfo", getAccountInfo{
		AccessToken: c.Account.AccessToken,
		Fields:      fields,
	})
	if err != nil {
		return nil, err
	}

	result := new(Account)
	if err = parser.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}
