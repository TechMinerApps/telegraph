package telegraph_test

import (
	"testing"

	"github.com/TechMinerApps/telegraph"
	"github.com/stretchr/testify/assert"
)

func Test_client_CreatePage(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		c := telegraph.NewClient()
		content, err := c.ContentFormat("<p>Hello, world!</p>")
		assert.NoError(t, err)
		_, err = c.CreatePage(telegraph.Page{
			Title:      "Sample Page",
			AuthorName: "Anonymous",
			Content:    content,
		}, true)
		assert.Error(t, err)
	})
	t.Run("Valid", func(t *testing.T) {
		c, err := telegraph.NewClientWithToken("b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb")
		assert.NoError(t, err)
		content, err := c.ContentFormat("<p>Hello, world!</p>")
		assert.NoError(t, err)

		page, err := c.CreatePage(telegraph.Page{
			Title:      "Sample Page",
			AuthorName: "Anonymous",
			Content:    content,
		}, true)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.NotEmpty(t, page.URL)
	})
}
