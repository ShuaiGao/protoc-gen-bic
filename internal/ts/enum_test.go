package ts

import "testing"

const enumResult = `export enum EDeviceOs {
  none    = 0, // 占位用
  IOS     = 1, // ios
  Android = 2, // android
}`

func TestEnum_String(t *testing.T) {
	type fields struct {
		name     string
		leading  string
		trailing string
		maxLen   int
		fields   []*EnumField
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{name: "tt1", fields: fields{name: "DeviceOs", leading: "// 设备系统", maxLen: 8, fields: []*EnumField{
			{FieldName: "none", Trailing: "// 占位用", Number: "0"},
			{FieldName: "IOS", Trailing: "// ios", Number: "1"},
			{FieldName: "Android", Trailing: "// android", Number: "2", LastOne: true},
		}}, want: enumResult},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Enum{
				Name:     tt.fields.name,
				Leading:  tt.fields.leading,
				Trailing: tt.fields.trailing,
				MaxLen:   tt.fields.maxLen,
				Fields:   tt.fields.fields,
			}
			got, err := e.String()
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
