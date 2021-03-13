package telegraph_test

import (
	"testing"

	"github.com/TechMinerApps/telegraph"
	"github.com/stretchr/testify/assert"
)

func Test_client_GetPage(t *testing.T) {
	c := telegraph.NewClient()
	t.Run("Invalid", func(t *testing.T) {
		_, err := c.GetPage("wtf", true)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		page, err := c.GetPage("Sample-Page-12-15", true)
		assert.NoError(t, err)
		assert.NotNil(t, page)
	})
}
