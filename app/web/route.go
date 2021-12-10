package web

import (
	"github.com/reallovelei/ggg/app/web/module/demo"
	"github.com/reallovelei/ggg/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist", "./dist")
	demo.Register(r)
}
