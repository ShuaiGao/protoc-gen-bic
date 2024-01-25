# protoc-gen-bic

protoc-gen-bic是一款强大、易用、优雅的接口生成工具。它设计用于解决基于protobuf的http接口定义，自动生成客户端（ts,js）和服务端（go)接口代码和接口文档。

服务端基于gin web框架构建模板，通过读取Protobuf的message和service构建服务注册和调用。

## 安装与验证

1. 安装

可在release中下载发布的 protoc-gen-bic 包

另外 protoc 文件下载地址 https://github.com/protocolbuffers/protobuf/releases/tag/v3.13.0


2. 验证

```shell
$ ./protoc-gen-bic --version
protoc-gen-bic 2.0.5
```

## 构建一个服务

服务构建，参见 example 目录下的 simple-go 项目，其中client目录仅仅用于存储客户端生成代码

或一个参照完整示例项目

- vue 前端，[bic-vue](https://github.com/ShuaiGao/bic-vue)
- go 后端，[bic-gin](https://github.com/ShuaiGao/bic-gin)
- proto，[bic-proto](https://github.com/ShuaiGao/bic-proto)

为避免重复造轮子，该示例中引入了一些其他工具，快捷方便地构建系统代码。

## 支持语言

支持生成下面语言的接口代码

- go
- ts
- js

## 特性

service 支持标签 @root，可用于设置url的前缀

在service定义中，rpc 方法通过注释指定api接口特性，特性支持下面标签(不区分大小写)

| 序号 | 标签名        | 标签含义                                                                                     |
|----|------------|------------------------------------------------------------------------------------------|
| 1  | Summary    | 接口注释                                                                                     |
| 2  | Doc        | 同 summary 标签                                                                             |
| 2  | Url        | 支持url参数，支持数据类型 int,uint,string,int64,uint64                                              |
| 3  | Method     | http方法，支持：get、post、patch、delete，可缺省使用，会自动识别pb定义的rpc方法名前缀                                 |
| 4  | Client-xxx | 用于生成客户端request填充自定义参数xxx                                                                 |
| 5  | void       | 服务端生成接口不返回数据，交由接口自行处理                                                                    |
| 6  | bind       | 服务端数据绑定，默认根据 method 区分使用 ShouldBindJSON、ShouldBindQuery 方法，当添加 bind 标签时，使用 ShouldBind 方法 |


## 联系

欢迎交流技术方法，邮箱：boringmanman@qq.com

<img src="https://tyimage.tuyoo.com/8751a21462/gaoshuai/wx.jpg" width="300px" height="auto">

## license

protoc-gen-bic is licensed under the [MIT](https://github.com/ShuaiGao/protoc-gen-bic/blob/main/LICENSE) license