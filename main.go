package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"protoc-gen-bic/internal"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), internal.VERSION)
		os.Exit(0)
	}
	if len(os.Args) == 2 && os.Args[1] == "--help" {
		os.Exit(0)
	}

	var (
		flags   flag.FlagSet
		plugins = flags.String("plugins", "", "deprecated option")
	)

	s := flags.String("ts_dir", "", "ts文件生成目录")
	responsePKG := flags.String("response_pkg", "", "返回消息通用消息头")
	permissionPKG := flags.String("permission_pkg", "", "权限校验包名")
	flag.Parse()
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		if *plugins != "" {
			return errors.New("protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC\n\n")
		}
		for _, f := range gen.Files {
			if f.Generate {
				if len(*responsePKG) > 0 {
					internal.ResponsePkg = *responsePKG
				}
				if len(*permissionPKG) > 0 {
					internal.PermissionPkg = *permissionPKG
				}
				internal.GenerateFile(gen, f)
				if len(*s) > 0 {
					internal.GenerateTsFile(gen, f, *s)
				}
			}
		}
		gen.SupportedFeatures = internal.SupportedFeatures
		return nil
	})
}
