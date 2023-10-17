package ts

import (
	"fmt"
	"strings"
	"text/template"
)

type Enum struct {
	Name     string
	Leading  string
	Trailing string
	MaxLen   int
	Fields   []*EnumField
}

func (e *Enum) AddField(field *EnumField) {
	fLen := field.GetLen()
	if fLen > e.MaxLen {
		e.MaxLen = fLen
	}
	e.Fields = append(e.Fields, field)
}

func (e *Enum) String() (string, error) {
	temp, err := template.ParseFiles("enum.tmpl")
	if err != nil {
		return "", nil
	}
	for _, v := range e.Fields {
		space := ""
		for i := 0; i < e.MaxLen-v.GetLen(); i++ {
			space += " "
		}
		v.Space = space
	}
	var buf strings.Builder
	err = temp.Execute(&buf, e)
	return buf.String(), err
}

type EnumField struct {
	FieldName string
	Number    string
	Trailing  string
	LastOne   bool
	Space     string
}

func (te *EnumField) GetLen() int {
	return len(te.FieldName) + len(te.Number)
}
func (te *EnumField) String(maxLen int) string {
	space := ""
	for i := 0; i < maxLen-te.GetLen(); i++ {
		space += " "
	}
	if te.LastOne {
		return fmt.Sprintf("  %s = %s  %s%s", te.FieldName, te.Number, space, te.Trailing)
	}
	return fmt.Sprintf("  %s = %s, %s%s", te.FieldName, te.Number, space, te.Trailing)
}
