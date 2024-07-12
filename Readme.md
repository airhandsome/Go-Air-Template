# Go Air Template代码结构框架

## 简介

这个框架旨在帮助用户快速创建一个项目框架。它集成了多种工具和库，如 `protoc-gen-grpc-gateway`、`protoc-gen-swagger`、`XO`、`Viper` 和 `Cobra`，以提供完整的开发体验。

## 目录说明

- `docs`: 文档文件夹，包含项目的安装指南、使用说明等。
- `src`: 源代码文件夹，包含Go语言编写的应用程序代码。
- `internal`: 内部模块文件夹，包含应用的核心逻辑和数据处理。
- `pkg`: 包文件夹，包含第三方库和工具。
- `log`: 日志文件夹，包含应用程序的日志文件。
- `bin`: 可执行文件夹，包含编译后的应用程序。
- `config`: 配置文件夹，包含项目的配置文件。
- `test`: 测试文件夹，包含单元测试和集成测试。
- `Dockerfile`: Dockerfile，用于构建和运行Docker容器。
- `Makefile`: Makefile，用于自动化项目构建、测试、打包等任务。
- `go.mod`: Go模块依赖管理文件。
- `go.sum`: Go模块依赖验证文件。

```
    /go-air-project
    ├── Dockerfile
    ├── LICENSE
    ├── Makefile
    ├── Readme.md
    ├── api
    │   └── proto
    │       ├── message.proto
    │       └── service.proto
    ├── cmd
    │   └── app
    │       ├── cmd.go
    │       └── main.go
    ├── config
    │   └── config.yaml
    ├── external
    │   └── google
    │       ├── api
    │       │   ├── annotations.proto
    │       │   └── http.proto
    │       └── protobuf
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── config
    │   ├── domain
    │   ├── handler
    │   ├── middleware
    │   ├── repository
    │   └── service
    ├── log
    │   └── log.go
    ├── scripts
    │   └── generate.sh
    └── sql
        ├── mongodb
        ├── psql
        │   └── example.sql
        └── redis
```

## 框架组成部分

1. **Protocol Buffers (proto)**: 用于定义RESTful API的HTTP路径和参数。
2. **Go语言**: 用于编写后端服务。
3. **Protoc-gen-grpc-gateway**: 用于生成gRPC到HTTP的网关代码。
4. **Protoc-gen-swagger**: 用于生成Swagger API文档。
5. **XO**: 用于从SQL文件生成Go代码。
6. **Viper**: 用于读取配置文件。
7. **Cobra**: 用于构建命令行界面。

## 安装和使用

1. **安装依赖**: 使用 `go get` 命令安装必要的库和工具。
2. **运行main**: 运行 `make run` 命令以启动应用。
3. **生成代码**: 使用 `make generate` 命令来生成Protocol Buffers、Swagger和Go代码。
4. **创建数据库模型**: 使用 `make sql-xo` 命令来生成数据库模型和操作函数。
5. **代码风格检查**: 使用 `make lint` 命令来检查代码风格。

## 注意事项

- 确保你的环境变量设置正确，如 `GOPATH` 和 `PATH`。
- 确保你的项目在模块模式下，并且 `go.mod` 文件包含了必要的依赖。
- 如果你遇到任何问题，请查看框架的README文件和源代码，或者在GitHub上创建一个Issue。

## 许可证

本框架遵循MIT许可证。
