package main

import (
	"educaition/dao"
	"educaition/routers"
)

func main() {
	// 初始化db
	dao.Init()
	r := routers.Router()
	r.Run(":8089")
}
