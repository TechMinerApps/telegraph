package telegraph_test

import (
	"testing"

	"github.com/TechMinerApps/telegraph"
	"github.com/valyala/fasthttp"
)

func Test_NewClientWithToken(t *testing.T) {

	t.Run("Valid", func(t *testing.T) {
		_, err := telegraph.NewClientWithToken("b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb", &fasthttp.Client{})
		if err != nil {
			t.FailNow()
		}
	})
	t.Run("Invalid", func(t *testing.T) {
		_, err := telegraph.NewClientWithToken("b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb", &fasthttp.Client{}, &fasthttp.Client{})
		if err == nil {
			t.FailNow()
		}
	})
}
