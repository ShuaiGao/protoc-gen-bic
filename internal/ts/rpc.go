package ts

import (
	"github.com/ShuaiGao/protoc-gen-bic/internal/utils"
	"strings"
	"text/template"
)

type Rpc struct {
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

func (r *Rpc) String() (string, error) {
	temp, err := template.ParseFiles("rpc.tmpl")
	if err != nil {
		return "", err
	}
	var buf strings.Builder
	if err = temp.Execute(&buf, r); err != nil {
		return "", err
	}
	return buf.String(), nil
}
