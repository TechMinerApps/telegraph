package telegraph_test

import (
	"testing"

	"github.com/TechMinerApps/telegraph"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	c := telegraph.NewClient()
	t.Run("Invalid", func(t *testing.T) {
		t.Run("Nil", func(t *testing.T) {
			_, err := c.CreateAccount(telegraph.Account{})
			assert.Error(t, err)
		})
		t.Run("Without shortname", func(t *testing.T) {
			_, err := c.CreateAccount(telegraph.Account{
				ShortName:  "",
				AuthorName: "Anonymous",
			})
			assert.Error(t, err)
		})
	})
	t.Run("valid", func(t *testing.T) {
		account, err := c.CreateAccount(telegraph.Account{
			ShortName:  "Sandbox",
			AuthorName: "Anonymous",
		})
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}
