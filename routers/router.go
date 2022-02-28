package routers

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()

	user(router)
	return router
}
