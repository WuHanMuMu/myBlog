package router

import (
	"github.com/gin-gonic/gin"
	"myBlog/control"
)

func R (e * gin.Engine) {
	// 前端文件
	e.Static("/static","./front")
	e.GET("/ping",control.Ping)
	articleRouter(e)
	userRouter(e)
	commentRouter(e)
}


func articleRouter (e *gin.Engine) {
	api := e.Group("/article")
	{
		api.POST("/Add",control.ArticleAdd)
		api.POST("/Detail",control.ArticleDetail)
		api.POST("/Pages",control.ArticlePage)
	}
}

func userRouter(e *gin.Engine){
	api := e.Group("/user")
	{
		api.POST("/Add",control.UserAdd)
		api.POST("/Detail",control.UserDetail)
	}
}

func commentRouter(e *gin.Engine){
	api := e.Group("/comment")
	{
		api.POST("/Add",control.CommentAdd)
		api.POST("/Detail",control.CommentDetail)
	}
}