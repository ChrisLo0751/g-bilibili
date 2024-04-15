package router

import (
	"g-shopping/user-web/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup, userHandler *api.UserHandler) {
	userRouter := router.Group("user")
	userRouter.GET("/:id", userHandler.GetUser)
}
