package initialize

import (
	"g-shopping/user-web/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()

	group := r.Group("/v1")
	router.InitUserRouter(group)

	return r
}
