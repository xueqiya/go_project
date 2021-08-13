package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xueqiya/go_project/controller"

	"github.com/xueqiya/go_project/config"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(config.Cfg.Server.Mode)
	apiv1 := r.Group("/api/v1")
	{
		// tags 路由
		apiv1.POST("/tags", controller.AddTag)
		apiv1.GET("/tags", controller.GetTags)
		apiv1.GET("/tags/:id", controller.GetTag)
		apiv1.PUT("/tags/:id", controller.EditTag)
		apiv1.DELETE("/tags/:id", controller.DeleteTag)

		// users 路由
		apiv1.POST("/users", controller.AddUser)
		apiv1.GET("/users/:id", controller.GetUser)
		apiv1.PATCH("/users/:id", controller.EditUser)
		apiv1.POST("/users/login", controller.Login)
	}
	return r
}
