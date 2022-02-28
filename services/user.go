package services

import (
	"educaition/dao"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserDao *dao.UserDao
}

func (u *User) Login(ctx *gin.Context) {
	//userName := ctx.PostForm("user_name")
	//password := ctx.PostForm("password")

}
