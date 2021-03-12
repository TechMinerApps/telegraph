package telegraph_test

import (
	"testing"

	"github.com/TechMinerApps/telegraph"
)

func Test_client_ContentFormat(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Invalid",
			args: args{
				data: 42,
			},
			wantErr: true,
		},
		{
			name: "String",
			args: args{
				data: "<p>Hello, World!</p>",
			},
			wantErr: false,
		},
		{
			name: "String",
			args: args{
				data: []byte("<p>Hello, World!</p>"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := telegraph.NewClient()
			gotN, err := c.ContentFormat(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.ContentFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (gotN == nil) != tt.wantErr {
				t.Errorf("client.ContentFormat() returns nil")
			}
		})
	}
}
