package initialize

import (
	"g-shopping/user-web/api"
	"g-shopping/user-web/data"
	"g-shopping/user-web/router"
	"g-shopping/user-web/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routers(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userData := data.NewUserData(db)
	userService := service.NewUserService(userData)
	userHandler := api.NewUserHandler(userService)

	group := r.Group("/v1")
	router.InitUserRouter(group, userHandler)

	return r
}
