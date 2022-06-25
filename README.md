# 说明

## 技术

1. 前端 vue
2. 后端 go
3. 数据库 mongodb

## go 初始化项目

1. 创建go.mod文件
```
go mod init word-dect-go
```
2. 引入gin框架
配置代理  
```
go env -w GOPROXY=https://goproxy.cn
```
```
go get -u github.com/gin-gonic/gin
```