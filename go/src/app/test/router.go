package test

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/test1", test1Handler)
	e.GET("/tesd2", test2Handler)
}
