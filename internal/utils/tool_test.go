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
