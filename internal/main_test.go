package internal

import (
	"reflect"
	"testing"
)

func Test_parseRpcLeading(t *testing.T) {
	type args struct {
		comm     string
		funcName string
	}
	tests := []struct {
		name      string
		args      args
		wantParam HTTPParam
	}{
		{
			name: "客户端参数",
			args: args{comm: "@cli-hello:true", funcName: "GetHello"},
			wantParam: HTTPParam{
				MethodName:      "GET",
				ClientParamList: []ClientParam{{Key: "hello", Value: "true"}},
			},
		},
		{
			name: "客户端参数1",
			args: args{comm: "@cli-hello :true @cli-world: 123", funcName: "GetHello"},
			wantParam: HTTPParam{
				MethodName: "GET",
				ClientParamList: []ClientParam{
					{Key: "hello", Value: "true"},
					{Key: "world", Value: "123"},
				},
			},
		},
		{
			name: "客户端参数2",
			args: args{comm: "@hello:world @cli-hello:true @cli-world:123 @cli-ddd:'uuu' @cli-sss:\"hello\"", funcName: "GetHello"},
			wantParam: HTTPParam{
				MethodName: "GET",
				ClientParamList: []ClientParam{
					{Key: "hello", Value: "true"},
					{Key: "world", Value: "123"},
					{Key: "ddd", Value: "'uuu'"},
					{Key: "sss", Value: "\"hello\""},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotParam := parseRpcLeading(tt.args.comm, tt.args.funcName); !reflect.DeepEqual(gotParam, tt.wantParam) {
				t.Errorf("parseRpcLeading() = %v, want %v", gotParam, tt.wantParam)
			}
		})
	}
}
