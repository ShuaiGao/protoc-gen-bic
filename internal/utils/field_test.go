package utils

import (
	"reflect"
	"testing"
)

func Test_parseLeading(t *testing.T) {
	type args struct {
		leading string
	}
	tests := []struct {
		name string
		args args
		want *FieldParam
	}{
		{name: "oneof", args: args{leading: "//@gotags: validate:\"oneof=asc desc\";"}, want: &FieldParam{FRequired: true, FEnums: []string{"asc", "desc"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLeading(tt.args.leading); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLeading() = %v, want %v", got, tt.want)
			}
		})
	}
}
