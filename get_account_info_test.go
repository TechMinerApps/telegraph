package telegraph_test

import (
	"testing"

	"github.com/TechMinerApps/telegraph"
	"github.com/stretchr/testify/assert"
)

func Test_client_GetAccountInfo(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		c := telegraph.NewClient()
		_, err := c.GetAccountInfo()
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		c, err := telegraph.NewClientWithToken("b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb")
		assert.NoError(t, err)

		info, err := c.GetAccountInfo(telegraph.FieldShortName, telegraph.FieldPageCount)
		assert.NoError(t, err)
		assert.Equal(t, "Sandbox", info.ShortName)
		assert.NotZero(t, info.PageCount)
	})
}
