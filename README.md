# oapi-codegen-example

## 运行、检查文档

1. `go run main.go`
1. 打开 localhost:8080/v1/apidocs

## 例子的实现步骤

1. 编辑 `spec/oapi/main.yaml`
1. `go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest`
1. `make genswagger`
1. 到 `app/app.go` 实现 Handler
1. `go run main.go`
1. 打开 localhost:8080/v1/apidocs
