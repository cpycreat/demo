package main

import (
	user "dydemo/api/user"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	//basegit push -u origin master
	//	apiRouter.POST("/user/login", user.Login)
	apiRouter.POST("/user/register", user.Register)
}
