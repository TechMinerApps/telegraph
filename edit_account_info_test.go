package telegraph_test

import (
	"testing"

	"github.com/TechMinerApps/telegraph"
	"github.com/stretchr/testify/assert"
)

func Test_client_EditAccountInfo(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		c := telegraph.NewClient()
		_, err := c.EditAccountInfo(telegraph.Account{})
		assert.Error(t, err)
	})
	t.Run("Valid", func(t *testing.T) {
		c, err := telegraph.NewClientWithToken("b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb")
		assert.NoError(t, err)

		_, err = c.EditAccountInfo(telegraph.Account{
			ShortName:  "Sandbox",
			AuthorName: "Anonymous",
		})
		assert.NoError(t, err)
	})
}
