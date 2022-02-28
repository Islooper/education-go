package services

import (
	"educaition/dao"
	"educaition/util"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserDao *dao.UserDao
}

func (u *User) Login(ctx *gin.Context) {
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")

	if userName == "" || password == "" {
		util.FailWithMessage("用户名或者密码不能为空", ctx)
		return
	}

	userDo, err := u.UserDao.ReadByUserNameAndPass(userName, password)
	if err != nil {
		util.FailWithMessage(err.Error(), ctx)
		return
	}

	if userDo.Id <= 0 {
		util.FailWithMessage("未找到该用户", ctx)
		return
	}

	util.OkWithData(userDo, ctx)

}
