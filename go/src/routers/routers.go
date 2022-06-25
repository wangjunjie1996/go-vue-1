package routers

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

// 注册 app 的路由配置, 吧 app\router.go 的 Routers func 放进 options 数组
func Include(opts []Option) {
	options = opts
}

// 初始化
func Init() *gin.Engine {
	r := gin.New()
	// 遍历每一个 Routers func 执行方法里的 e.GET 等
	for _, opt := range options {
		opt(r)
	}
	return r
}
