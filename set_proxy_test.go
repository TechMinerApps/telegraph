package telegraph_test

import (
	"testing"

	"github.com/TechMinerApps/telegraph"
)

func Test_client_SetSocksProxy(t *testing.T) {

	tests := []struct {
		name string
	}{
		{name: "Empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := telegraph.NewClient()
			c.SetSocksProxy("")
		})
	}
}

func Test_client_SetHTTPProxy(t *testing.T) {

	tests := []struct {
		name string
	}{
		{name: "Empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := telegraph.NewClient()
			c.SetHTTPProxy("")
		})
	}
}
