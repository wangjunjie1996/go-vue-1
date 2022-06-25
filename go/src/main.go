package main

import (
	"fmt"
	"my-go1/src/app/hello"
	"my-go1/src/app/test"
	"my-go1/src/routers"
)

var routs = []routers.Option{
	hello.Routers,
	test.Routers,
}

func main() {
	// 加载多个APP的路由配置
	routers.Include(routs)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
