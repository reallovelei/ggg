package web

import (
	"github.com/reallovelei/ggg/app/web/module/demo"
	"github.com/reallovelei/ggg/framework/gin"
	"github.com/reallovelei/ggg/framework/middleware/static"
)

func Routes(r *gin.Engine) {
	// /路径先去./dist目录下查找文件是否存在，如果找到, 使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./front/dist", false)))
	//	r.Static("/dist", "./dist")
	demo.Register(r)
}
