package ts

import (
	"fmt"
	"strings"
	"text/template"
)

type Interface struct {
	Name     string
	Leading  string
	Trailing string
	MaxLen   int
	Fields   []*InterfaceField
}

func (i *Interface) String() (string, error) {
	temp, err := template.ParseFiles("interface.tmpl")
	if err != nil {
		return "", err
	}
	//for _, v := range i.Fields {
	//	space := ""
	//	for i := 0; i < i.MaxLen-v.GetLen(); i++ {
	//		space += " "
	//	}
	//	v.Space = space
	//}
	var buf strings.Builder
	if err = temp.Execute(&buf, i); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type InterfaceField struct {
	FieldName string
	FieldType string
	Trailing  string
	Required  bool
	IsList    bool
	TotalLen  int
}

func (tf *InterfaceField) String(maxLen int) string {
	space := ""
	for i := 0; i < maxLen-tf.TotalLen; i++ {
		space += " "
	}
	return fmt.Sprintf("  %s: %s %s%s", tf.FieldName, tf.FieldType, space, tf.Trailing)
}
