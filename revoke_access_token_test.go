package telegraph_test

import (
	"testing"

	"github.com/TechMinerApps/telegraph"
	"github.com/stretchr/testify/assert"
)

func Test_client_RevokeAccessToken(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		c := telegraph.NewClient()
		_, err := c.RevokeAccessToken()
		assert.Error(t, err)
	})
	t.Run("Valid", func(t *testing.T) {
		c := telegraph.NewClient()
		_, err := c.CreateAccount(telegraph.Account{
			ShortName:  "Sandbox",
			AuthorName: "Anonymous",
		})
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		newAccount, err := c.RevokeAccessToken()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.NotEqual(t, c.Account().AccessToken, newAccount.AccessToken)
		assert.NotEmpty(t, newAccount.AuthURL)
	})
}
