# oapi-codegen-example

## 例子的实现步骤

1. 编辑 `spec/oapi/main.yaml`
2. `make genswagger`
3. 到 `app/app.go` 实现 Handler
4. `go run main.go`
5. 打开 localhost:8080/v1/apidocs
