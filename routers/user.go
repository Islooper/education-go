package routers

import (
	"educaition/dao"
	"educaition/services"
	"github.com/gin-gonic/gin"
)

func user(ctx *gin.Engine) {
	user := &services.User{
		UserDao: new(dao.UserDao),
	}
	ctx.POST("/user/login", user.Login)

}
