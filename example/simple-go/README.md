bic 使用示例

本示例，基于gin web框架

本示例将展示通过protobuf定义生成gin接口interface和接口swagger注释，并自动生成swagger注释文档。具体命令可参照`makefile`文件

# 生成命令

1. 前端命令 `make ts`
2. 后端命令 `make pb`

# 安装依赖程序
## 后端依赖程序

1. 安装 protoc
    下载地址 https://github.com/protocolbuffers/protobuf/releases/

2. 安装protoc-gen-bic

```shell
go install github.com/ShuaiGao/protoc-gen-bic@latest
```

3. 安装 protoc-gen-go

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

4. 安装 protoc-go-inject-tag

```shell
go install github.com/favadi/protoc-go-inject-tag@latest
```

5. 安装 swag

```shell
go install github.com/swaggo/swag/cmd/swag
```

## 前端依赖程序

1. 安装pbjs `npm install pgjs -g`
2. 安装pbts `npm install pgts -g`
