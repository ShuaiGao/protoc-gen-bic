# protoc-gen-bic

protoc-gen-bic是一款强大、易用、优雅的接口生成工具。它设计用于解决基于protobuf的http接口定义，自动生成客户端（ts）和服务端（go)接口代码和接口文档。

服务端基于gin web框架构建模板，通过读取Protobuf的message和service构建服务注册和调用。

## 安装与验证

1. 安装

```git
go install github.com/ShuaiGao/protoc-gen-bic@latest
```

2. 验证

> protoc-gen-bic --version
> 
> protoc-gen-bic v1.0.0

总结: 下载代码目的是将二进制文件下载到gopath的bin文件下。

## 构建一个服务

服务构建，参见 example 目录下的 simple-go 项目，其中client目录仅仅用于存储客户端生成代码

为避免重复造轮子，该示例中引入了一些其他工具，快捷方便地构建系统代码。

## 支持语言

支持生成下面语言的接口代码

- go
- ts
- js

## 特性

在service定义中，通过注释指定api接口特性，特性支持下面标签(不区分大小写)

- Summary
- Url，支持url参数，支持数据类型 int,uint,string,int64,uint64
- Method
- Client-xxx 用于生成客户端request填充自定义参数xxx

## license

protoc-gen-bic is licensed under the [MIT](https://github.com/ShuaiGao/protoc-gen-bic/blob/main/LICENSE) license