package v1

//添加protoc接口协议文件
//kratos proto add api/helloworld/demo.proto

//基于proto文件生成客户端代码
//kartos proto client api/helloworld/v1/greeter.proto

//通过 proto文件，生成对应的 Service 实现代码：
//kratos proto server api/helloworld/demo.proto -t internal/service

//运行项目
//kratos run

//wire 代码生成
//go generate ./...
