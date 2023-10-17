package ts

import (
	"github.com/ShuaiGao/protoc-gen-bic/internal/utils"
	"testing"
)

func TestRpc_String(t *testing.T) {
	type fields struct {
		Name     string
		Url      string
		Input    string
		Output   string
		Leading  string
		Trailing string
		MaxLen   int
		Method   string
		params   []*utils.URLParam
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{name: "tt1", fields: fields{
			Name:    "GetRequestUsers",
			Url:     "/api/v1/users",
			Leading: "// 测试接口",
			Method:  "GET", Input: "RequestUsers", Output: "ResponseUsers", MaxLen: 8},
			want: interfaceResult, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rpc{
				Name:     tt.fields.Name,
				Url:      tt.fields.Url,
				Input:    tt.fields.Input,
				Output:   tt.fields.Output,
				Leading:  tt.fields.Leading,
				Trailing: tt.fields.Trailing,
				MaxLen:   tt.fields.MaxLen,
				Method:   tt.fields.Method,
				params:   tt.fields.params,
			}
			got, err := r.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("String() got = %v, want %v", got, tt.want)
			}
		})
	}
}
