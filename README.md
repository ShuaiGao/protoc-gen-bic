# 自动生成代码介绍

> protoc-gen-bic 是一款可以自动生成gin http服务注册和调用的工具，其基于gin web框架构建模板，通过读取Protobuf的message和service构建服务注册和调用

## 构建生成代码工具

1. 安装
```git
// go1.17前
go get -u github.com/ShuaiGao/protoc-gen-bic@latest
go install github.com/ShuaiGao/protoc-gen-bic
// go1.17后
go install github.com/ShuaiGao/protoc-gen-bic@latest
```
2. 验证
> protoc-gen-bic --version
> 
> protoc-gen-bic v1.0.0

总结: 下载代码目的是将二进制文件下载到gopath的bin文件下。

## 构建一个服务

参见示例

> 所有的自定义代码都是通过注释读取的。

- 自定义组中间件
> 每一个service对应gin都是一个组，定义组路由只需要使用`@middle:`关键字即可，value是对应的方法
- 导入项目包
> 在service中定义 `@import` 可以连续定义多个导入包

- 定义接口请求

> 在每个rpc接口中,自定义请求方法 `@method` 即可目前支持GET，POST，ANY 不区分大小写

- 定义接口中间件
> 在接口注释中 `@middle` 即可定义自定义和gin中间件

- 如何自定义组路由

> 在service 定义`@root`即可

