# Go Gin 文档   1.7 版本

## Gin 是什么？

Gin 是一个用 Go (Golang) 编写的 HTTP web 框架。 它是一个类似于 martini 但拥有更好性能的 API 框架，由于 httprouter，速度提高了近 40 倍。如果你需要极好的性能，使用 Gin 吧。

## 如何使用 Gin？

我们提供了一些 API 的 使用示例 并罗列了一些众所周知的 Gin 使用者.

## 如何为 Gin 做贡献？

    在社区帮助其他人
    告诉我们你使用 Gin 的成功案例
    告诉我们如何改善 Gin 并帮助我们实现它
    为已有的库做贡献

# Gin 简介 

Gin 是一个用 Go (Golang) 编写的 web 框架。 它是一个类似于 martini 但拥有更好性能的 API 框架，由于 httprouter，速度提高了近 40 倍。 如果你是性能和高效的追求者，你会爱上 Gin.

在本节中，我们将介绍 Gin 是什么，它解决了哪些问题，以及它如何帮助你的项目。

或者，如果你已经准备在项目中使用 Gin，请访问快速入门.

## 特性

### 快速

基于 Radix 树的路由，小内存占用。没有反射。可预测的 API 性能。

### 支持中间件

传入的 HTTP 请求可以由一系列中间件和最终操作来处理。 例如：Logger，Authorization，GZIP，最终操作 DB。

### Crash 处理

Gin 可以 catch 一个发生在 HTTP 请求中的 panic 并 recover 它。这样，你的服务器将始终可用。例如，你可以向 Sentry 报告这个 panic！

### JSON 验证

Gin 可以解析并验证请求的 JSON，例如检查所需值的存在。

### 路由组

更好地组织路由。是否需要授权，不同的 API 版本…… 此外，这些组可以无限制地嵌套而不会降低性能。

### 错误管理

Gin 提供了一种方便的方法来收集 HTTP 请求期间发生的所有错误。最终，中间件可以将它们写入日志文件，数据库并通过网络发送。

### 内置渲染

Gin 为 JSON，XML 和 HTML 渲染提供了易于使用的 API。

### 可扩展性

新建一个中间件非常简单，去查看[示例代码](https://gin-gonic.com/zh-cn/docs/examples/using-middleware/)吧。

# 快速入门

## 要求

* Go 1.13 及以上版本 

## 安装

要安装 Gin 软件包，需要先安装 Go。

1. 下载并安装 gin：
```
$ go get -u github.com/gin-gonic/gin
```

2. 将 gin 引入到代码中：
```
import "github.com/gin-gonic/gin"
```
3.（可选）如果使用诸如 http.StatusOK 之类的常量，则需要引入 net/http 包：
```
import "net/http"
```
运行示例

1. 创建项目并且 cd 到项目目录中 (也可以把项目放到你想放的任何地方，使用 go mod 的好处就是不必强制把项目放到 GOPATH 下了)
```
$ mkdir -p $GOPATH/src/github.com/myusername/project && cd "$_"
```
2. 初始化 go mod
```
$ go mod init project
```
3. 启动项目
```
$ go run main.go
```
开始
```
    不确定如何编写和执行 Go 代码？点击这里
```
首先，创建一个名为 example.go 的文件：
```
$ touch example.go
```
接下来，将如下的代码写入 example.go 中：
```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
```
然后，执行 go run example.go 命令来运行代码：

### 运行 example.go 并且在浏览器中访问 0.0.0.0:8080/ping
```
$ go run example.go
```

# Jsoniter

## 使用 jsoniter 编译

Gin 使用 encoding/json 作为默认的 json 包，但是你可以在编译中使用标签将其修改为 jsoniter。
```
$ go build -tags=jsoniter .
```

# 部署

Gin 项目可以很方便地在以下这些云平台上部署：

## Qovery

Qovery 云平台提供免费的数据库、SSL 和一个全球 CDN，并支持使用 Git 自动部署。

详见 [部署你的 Gin 项目.](https://docs.qovery.com/guides/tutorial/deploy-gin-with-postgresql/)
## Render

Render 原生支持 Go，另外提供了 SSL 管理、数据库、平滑部署、 HTTP/2、 websocket 等支持。文档 [如何部署 Gin 项目到 Render 上.](https://render.com/docs/deploy-go-gin)

## Google App Engine

GAE 有两个部署 Go 的方式。一个是标准环境，比较容易使用，但是定制起来比较不方便，并为安全提供了 [syscalls](https://github.com/gin-gonic/gin/issues/1639)。另一个是灵活的环境，可以运行各种框架和类库。

如何选择请查看文档： [Google App Engine 的 Go 应用.](https://cloud.google.com/appengine/docs/go/)

# 测试

## 怎样编写 Gin 的测试用例

HTTP 测试首选 net/http/httptest 包。
```go
package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
    return r
}

func main() {
    r := setupRouter()
    r.Run(":8080")
}
```
上面这段代码的测试用例：
```go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/ping", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Equal(t, "pong", w.Body.String())
}
```