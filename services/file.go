package services

import (
	"educaition/dao"
	"educaition/util"
	"errors"
	"github.com/gin-gonic/gin"
	"os"
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

func (f *File) Delete(ctx *gin.Context) {
	id := ctx.GetInt("id")
	if id == 0 {
		util.FailWithMessage(errors.New("参数不能为空").Error(), ctx)
		return
	}

	// 查找文件
	fileDo, err := f.FileDao.ReadById(int64(id))
	if err != nil {
		util.FailWithMessage(errors.New("查找文件失败").Error(), ctx)
		return
	}

	if fileDo.ID <= 0 {
		util.FailWithMessage(errors.New("文件不存在").Error(), ctx)
		return
	}

	// 删除数据库记录
	err = f.FileDao.Delete(int64(id))
	if err != nil {
		util.FailWithMessage(errors.New("删除失败").Error(), ctx)
		return
	}

	// 删除本地文件
	err = os.Remove(fileDo.Path)
	if err != nil {
		util.FailWithMessage(errors.New("删除失败").Error(), ctx)
		return
	}

	util.Ok(ctx)
}

func (f *File) AgreeOrRefuse(ctx *gin.Context) {
	id := ctx.GetInt("id")
	param := ctx.GetInt("param")

	if id == 0 || param == 0 {
		util.FailWithMessage(errors.New("参数不能为空").Error(), ctx)
		return
	}

	// 找到该文件
	fileDo, err := f.FileDao.ReadById(int64(id))
	if err != nil {
		util.FailWithMessage(errors.New("查找文件失败").Error(), ctx)
		return
	}

	// 更新文件
	err = f.FileDao.UpdateIsOn(int64(fileDo.ID), param)
	if err != nil {
		util.FailWithMessage(errors.New("更新失败").Error(), ctx)
		return
	}

	util.Ok(ctx)

}
