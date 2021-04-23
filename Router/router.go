package Router

import (
	"gin-demo/Controllers/user"
	"gin-demo/Middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	//全局使用跨域中间件
	router.Use(Middlewares.Cors())

	//路由分组
	api1 := router.Group("api1")
	{
		api1.POST("/login", user.Login)
		api1.POST("/register", user.Register)
	}

	router.Run(":8088")
}
