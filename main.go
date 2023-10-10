package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/ShuaiGao/protoc-gen-bic/internal"
	"github.com/ShuaiGao/protoc-gen-bic/internal/js"
	"github.com/ShuaiGao/protoc-gen-bic/internal/ts"
	"github.com/ShuaiGao/protoc-gen-bic/internal/utils"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), utils.VERSION)
		os.Exit(0)
	}
	if len(os.Args) == 2 && os.Args[1] == "--help" {
		os.Exit(0)
	}

	var (
		flags   flag.FlagSet
		plugins = flags.String("plugins", "", "deprecated option")
	)

	tsDir := flags.String("ts_dir", "", "ts文件生成目录")
	jsDir := flags.String("js_dir", "", "js文件生成目录")
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
				if len(*permissionPKG) > 0 {
					internal.PermissionPkg = *permissionPKG
				}
				if len(*tsDir) > 0 {
					ts.GenerateTsFile(gen, f, *tsDir)
				} else if len(*jsDir) > 0 {
					js.GenerateJsFile(gen, f, *jsDir)
				} else {
					internal.GenerateFile(gen, f)
				}
			}
		}
		gen.SupportedFeatures = internal.SupportedFeatures
		return nil
	})
}
