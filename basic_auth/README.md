# Basic Auth
前往仓库根目录，启动服务器
```shell
$ go run basic_auth/main.go
```
进行测试

1. 直接用浏览器访问 `http://localhost:5444`，浏览器在收到响应头 `WWW-Authenticate: realm="protected"` 和状态码 `401 Unauthorized` 后会弹出登录框，这是浏览器实现的
2. 用浏览器访问 `http://Foo:Bar@localhost:5444`，可以直接看到 `Authorized` 字样
3. 用 `Postman` 请求 `GET http://localhost:5444`，并提供请求头 `Authorization: Basic Rm9vOkJhcg==` 也能收到 `Authorized`
