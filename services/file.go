package services

import (
	"educaition/dao"
	"educaition/util"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type File struct {
	FileDao *dao.FileDao
}

func (f *File) List(ctx *gin.Context) {
	pageNo := ctx.GetInt("page_no")
	pageSize := ctx.GetInt("page_size")

	fileDos, err := f.FileDao.List(pageNo, pageSize)
	if err != nil {
		util.FailWithMessage(err.Error(), ctx)
		return
	}

	util.OkWithData(fileDos, ctx)
}

func (f *File) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	userId := ctx.PostForm("user_id")
	form := ctx.PostForm("form")

	if err != nil {
		util.FailWithMessage(err.Error(), ctx)
		return
	}

	if file.Size == 0 {
		util.FailWithMessage(errors.New("文件为空").Error(), ctx)
		return
	}
	//将文件保存至本项目根目录中
	err = ctx.SaveUploadedFile(file, file.Filename)
	if err != nil {
		util.FailWithMessage(err.Error(), ctx)
		return
	}

	uId, err := strconv.ParseInt(userId, 10, 64)
	fileDo := &dao.File{
		Name:   file.Filename,
		Size:   file.Size,
		Form:   form,
		Path:   file.Filename,
		UserId: uId,
	}

	err = f.FileDao.Create(fileDo)
	if err != nil {
		util.FailWithMessage(err.Error(), ctx)
		return
	}

	util.Ok(ctx)
}
