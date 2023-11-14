package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Max[T int | uint | int64 | uint64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type URLParam struct {
	PType string
	PName string
}

type ClientParam struct {
	Key   string
	Value string
}
type HTTPParam struct {
	Summary         string
	MethodName      string
	Url             string
	UrlParamList    []URLParam
	Permission      string
	Void            bool
	ClientParamList []ClientParam
	Download        bool
}

// ParseRpcLeading 解析rpc方法注释
func ParseRpcLeading(comm string, funcName string) (param HTTPParam) {
	param = HTTPParam{}
	supportMethods := []string{"GET", "POST", "DELETE", "PATCH"}

	for _, m := range supportMethods {
		// 匹配方法命名规则
		matched, err := regexp.MatchString(`(?i)`+m+`\w+`, funcName)
		if err == nil && matched {
			param.MethodName = m
			break
		}
	}
	// 优先使用 method标注
	for _, m := range supportMethods {
		matched, err := regexp.MatchString(`(?i)@method\s*:\s*`+m, comm)
		if err == nil && matched {
			param.MethodName = m
			break
		}
	}
	re := regexp.MustCompile(`(?i)@url\s*:\s*(/.*)\s`)
	urls := re.FindSubmatch([]byte(comm))
	if len(urls) > 1 {
		param.Url = strings.TrimSpace(string(urls[1]))
	}
	paramRe := regexp.MustCompile(`/<?(\w*):(\w+)>?`)
	matchList := paramRe.FindAllSubmatch([]byte(param.Url), -1)
	for _, match := range matchList {
		pName := string(match[2])
		pType := string(match[1])
		if pType == "" {
			pType = StringType
		} else if pType == "str" || pType == StringType {
			pType = StringType
		} else if pType != IntType && pType != Int64Type && pType != UIntType {
			fmt.Fprintf(os.Stderr, "path param only support int and string/str, but get "+pType)
			os.Exit(0)
		}
		param.UrlParamList = append(param.UrlParamList, URLParam{
			PName: pName,
			PType: pType,
		})
	}
	delTypeRe := regexp.MustCompile(`<(\w+):(\w+)>`)
	param.Url = delTypeRe.ReplaceAllString(param.Url, ":$2")

	rePermission := regexp.MustCompile(`(?i)@permission\s*:\s*(\w*)`)
	objs := rePermission.FindSubmatch([]byte(comm))
	if len(objs) > 1 {
		param.Permission = strings.TrimSpace(string(objs[1]))
	}

	summaryPermission := regexp.MustCompile(`(?i)@summary\s*:\s*([^@$]*)`)
	summary := summaryPermission.FindSubmatch([]byte(comm))
	if len(summary) > 1 {
		param.Summary = strings.TrimSpace(string(summary[1]))
	} else {
		regexpDoc := regexp.MustCompile(`(?i)@doc\s*:\s*([^@$]*)`)
		doc := regexpDoc.FindSubmatch([]byte(comm))
		if len(doc) > 1 {
			param.Summary = strings.TrimSpace(string(doc[1]))
		}
	}

	returnPermission := regexp.MustCompile(`(?i)@void`)
	void := returnPermission.FindSubmatch([]byte(comm))
	if len(void) > 0 {
		param.Void = true
	}

	downloadPermission := regexp.MustCompile(`(?i)@download`)
	download := downloadPermission.FindSubmatch([]byte(comm))
	if len(download) > 0 {
		param.Void = true
		param.Download = true
	}

	clientParam := regexp.MustCompile(`(?i)@cli-(\w+)\s*:\s*(["'a-zA-Z0-9]*)`)
	paramList := clientParam.FindAllSubmatch([]byte(comm), -1)
	for _, v := range paramList {
		if len(v) > 2 {
			param.ClientParamList = append(param.ClientParamList, ClientParam{
				Key:   string(v[1]),
				Value: string(v[2]),
			})
		}
	}
	return
}

type ServiceParam struct {
	Root string
}

func ParseServiceLeading(comments string) ServiceParam {
	re := regexp.MustCompile(`(?i)@root\s*:\s*([^@$]*)`)
	rootUrls := re.FindSubmatch([]byte(comments))

	var param ServiceParam
	if len(rootUrls) > 1 {
		param.Root = strings.TrimSpace(string(rootUrls[1]))
	}
	return param
}
