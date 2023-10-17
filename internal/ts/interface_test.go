package ts

import "testing"

const interfaceResult = `export interface IRequestUsers {
  page: number, // 页码
  page_size: number, // 每页数量
}`

func TestInterface_String1(t *testing.T) {
	type fields struct {
		Name     string
		Leading  string
		Trailing string
		MaxLen   int
		Fields   []*InterfaceField
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr error
	}{
		{name: "tt1", fields: fields{Name: "RequestUsers", Leading: "// 设备系统", MaxLen: 8, Fields: []*InterfaceField{
			{FieldType: "number", Trailing: "// 页码", Required: true, FieldName: "page"},
			{FieldType: "number", Trailing: "// 每页数量", Required: true, FieldName: "page_size"},
		}}, want: interfaceResult, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interface{
				Name:     tt.fields.Name,
				Leading:  tt.fields.Leading,
				Trailing: tt.fields.Trailing,
				MaxLen:   tt.fields.MaxLen,
				Fields:   tt.fields.Fields,
			}
			got, err := i.String()
			if tt.wantErr != err {
				t.Errorf("String() = %v, want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
