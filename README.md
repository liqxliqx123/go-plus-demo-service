# my-demo-SERVICE

## Code status


## 项目介绍

## 架构介绍

### gRPC

该脚手架集成了 `gRPC` 和 `grpc-gateway`。使用时，请注意：

* 在 `proto/service.proto` 中定义函数。
* 执行 `make protod` 命令生成相关代码。生成的文件有：
  * protobuf 文件
  * grpc 文件
  * grpc-gateway 文件
  * grpc mock 文件
  * swagger 文档
