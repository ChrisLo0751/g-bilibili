package router

import (
	"g-shopping/user-web/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	userRouter.GET("list", api.GetUserList)
}