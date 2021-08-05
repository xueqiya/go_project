package router

import (
	"github.com/gin-gonic/gin"

	"github.com/linehk/gin-blog/config"
	"github.com/linehk/gin-blog/router/api/v1"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(config.Cfg.Server.Mode)
	apiv1 := r.Group("/api/v1")
	{
		//// articles 路由
		//apiv1.GET("/articles", v1.GetArticles)
		//apiv1.GET("/articles/:id", v1.GetArticle)
		//apiv1.POST("/articles", v1.AddArticle)
		//apiv1.PUT("/articles/:id", v1.EditArticle)
		//apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//
		//// tags 路由
		//apiv1.GET("/tags", v1.GetTags)
		//apiv1.GET("/tags/:id", v1.GetTag)
		//apiv1.POST("/tags", v1.AddTag)
		//apiv1.PUT("/tags/:id", v1.EditTag)
		//apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// dynamics 路由
		apiv1.GET("/dynamics", v1.GetTags)
		apiv1.GET("/dynamics/:id", v1.GetTag)
		apiv1.POST("/dynamics", v1.AddTag)
		apiv1.PUT("/dynamics/:id", v1.EditTag)
		apiv1.DELETE("/dynamics/:id", v1.DeleteTag)
	}
	return r
}
