package Router

import (
	"demoTest/Controllers"
	"demoTest/Middlewares"
	"demoTest/Sessions"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//RESTful APi设计
//Get用来获取资源
//Post用来新建资源
//Put用来更新资源
//Delete用来删除资源
func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	// 使用 session(cookie-based)
	router.Use(sessions.Sessions("mysession", Sessions.Store))
	//对接口进行分类
	v1 := router.Group("/api/user")
	{
		v1.POST("/login", Controllers.Login)
	}
	v2 :=router.Group("/api/file")
	{
		v2.POST("/upload", Controllers.Upload)
		v2.GET("/getimage/:imageName",Controllers.GetImage)
	}

	router.Run(":8088")
}
