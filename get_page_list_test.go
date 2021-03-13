package telegraph_test

import (
	"testing"

	"github.com/TechMinerApps/telegraph"
	"github.com/stretchr/testify/assert"
)

func Test_client_GetPageList(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		c := telegraph.NewClient()
		_, err := c.GetPageList(0, 0)
		assert.Error(t, err)
	})
	t.Run("Valid", func(t *testing.T) {
		c, err := telegraph.NewClientWithToken("b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb")

		assert.NoError(t, err)

		list, err := c.GetPageList(1, 1)
		assert.NoError(t, err)
		assert.NotNil(t, list)
	})
}
