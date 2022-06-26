# 说明

## 技术

1. 前端 vue3
2. 后端 go
3. 数据库 mongodb

## go 初始化项目

1. 创建go.mod文件
```
go mod init word-dect-go
```
2. [引入gin框架](https://learnku.com/docs/gin-gonic/1.7/go-gin-document/11352)
配置代理  
```
go env -w GOPROXY=https://goproxy.cn
```
```
go get -u github.com/gin-gonic/gin
```

## 运行
```
go run main.go
```
http://localhost:8080/${api-url}