package utils

import (
	"reflect"
	"testing"
)

func TestParseServiceLeading(t *testing.T) {
	type args struct {
		comments string
	}
	tests := []struct {
		name string
		args args
		want ServiceParam
	}{
		{name: "root-1", args: args{comments: "// @root: ook"}, want: ServiceParam{Root: "ook"}},
		{name: "root-2", args: args{comments: "// @Root: ook"}, want: ServiceParam{Root: "ook"}},
		{name: "root-3", args: args{comments: "//@ROOT: /ook"}, want: ServiceParam{Root: "/ook"}},
		{name: "root-4", args: args{comments: "//@ROOT : /ook"}, want: ServiceParam{Root: "/ook"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseServiceLeading(tt.args.comments); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseServiceLeading() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseRpcLeading(t *testing.T) {
	type args struct {
		comm     string
		funcName string
	}
	tests := []struct {
		name      string
		args      args
		wantParam HTTPParam
	}{
		{name: "get", args: args{comm: "", funcName: "GetName"}, wantParam: HTTPParam{MethodName: "GET"}},
		{name: "uint", args: args{comm: "@url:/hello/<uint:id>/", funcName: "GetName"}, wantParam: HTTPParam{
			MethodName: "GET", Url: "/hello/:id/", UrlParamList: []URLParam{{PType: UIntType, PName: "id"}},
		}},
		{name: "uint", args: args{comm: "@url:/hello/<uint:id>", funcName: "GetName"}, wantParam: HTTPParam{
			MethodName: "GET", Url: "/hello/:id", UrlParamList: []URLParam{{PType: UIntType, PName: "id"}},
		}},
		{name: "uint", args: args{comm: "@url:/hello/<uint:id> @method:GET", funcName: "GetName"}, wantParam: HTTPParam{
			MethodName: "GET", Url: "/hello/:id", UrlParamList: []URLParam{{PType: UIntType, PName: "id"}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotParam := ParseRpcLeading(tt.args.comm, tt.args.funcName); !reflect.DeepEqual(gotParam, tt.wantParam) {
				t.Errorf("ParseRpcLeading() = %v, want %v", gotParam, tt.wantParam)
			}
		})
	}
}

func Test_genTsJsUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "url-1", args: args{url: "/hello/world"}, want: `"/hello/world"`},
		{name: "url-2", args: args{url: "/hello/:id"}, want: `"/hello/" + id`},
		{name: "url-3", args: args{url: "/hello/:id/world"}, want: `"/hello/" + id + "/world"`},
		{name: "url-4", args: args{url: "/hello/:id/world/"}, want: `"/hello/" + id + "/world/"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genTsJsUrl(tt.args.url); got != tt.want {
				t.Errorf("genTsJsUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
