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
		// users 路由
		apiv1.POST("/users", controller.AddUser)
		apiv1.GET("/users/:id", controller.GetUser)
		apiv1.PATCH("/users/:id", controller.EditUser)
		apiv1.POST("/users/login", controller.Login)

		// goods 路由
		apiv1.GET("/goods", controller.GetAllGoods)
		apiv1.POST("/goods", controller.AddGoods)
		apiv1.GET("/goods/:id", controller.GetGoods)
		//apiv1.PUT("/goods/:id", controller.EditGoods)
		//apiv1.DELETE("/goods/:id", controller.DeleteGoods)
	}
	return r
}
