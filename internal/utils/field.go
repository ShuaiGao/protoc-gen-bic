package utils

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"regexp"
	"strings"
)

type FieldParam struct {
	FType     string
	FName     string
	FTail     string
	FLeading  string
	FEnums    []string
	FMin      int
	FMax      int
	FRequired bool
}

func (fp *FieldParam) GetEnums() string {
	if len(fp.FEnums) == 0 {
		return ""
	}
	return fmt.Sprintf("Enums(%s)", strings.Join(fp.FEnums, ","))
}

// ParseFieldLeading 解析field注释
func ParseFieldLeading(field *protogen.Field) *FieldParam {
	leading := field.Comments.Leading.String()
	param := parseLeading(leading)
	param.FName = JSONSnakeCase(field.GoName)
	//param.FName = field.Desc.JSONName()
	if !param.FRequired {
		param.FRequired = strings.Contains(leading, Required)
	}

	trialing := field.Comments.Trailing.String()
	if len(trialing) > 2 {
		param.FTail = strings.TrimSpace(field.Comments.Trailing.String()[2:])
	}

	kindName := field.Desc.Kind().String()
	if field.Enum != nil {
		kindNames := strings.Split(string(field.Enum.Desc.FullName()), ".")
		if len(kindNames) < 3 {
			kindName = string(field.Enum.Desc.Name())
		} else {
			kindName = strings.Join(kindNames[1:], "_")
		}
	} else if field.Desc.IsList() && field.Desc.Message() != nil {
		kindName = fmt.Sprintf("[]%s", field.Desc.Message().Name())
	}
	if kindName == "message" {
		kindName = string(field.Desc.Message().Name())
		if kindName == "DataEntry" {
			// TODO 处理 map 类型
			kindName = "object"
		}
	} else if kindName == "bytes" {
		kindName = "[]byte"
	} else if kindName == "float" {
		kindName = "float32"
	} else if kindName == "double" {
		kindName = "float64"
	}
	if field.Desc.IsList() {
		if field.Desc.Enum() != nil {
			kindName = "int32"
		}
		if !strings.HasPrefix(kindName, "[]") {
			kindName = "[]" + kindName
		}
	} else if field.Desc.IsMap() {
		kindName = "object"
	}
	param.FType = kindName
	return param
}

func parseLeading(leading string) *FieldParam {
	param := &FieldParam{}
	re := regexp.MustCompile(`(?i)@gotags\s*:\s*validate:"oneof=(.*)";`)
	oneof := re.FindSubmatch([]byte(leading))
	if len(oneof) > 1 {
		param.FRequired = true
		param.FEnums = strings.Split(string(oneof[1]), " ")
	}
	return param
}
