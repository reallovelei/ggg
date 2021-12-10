package http

import (
	"github.com/reallovelei/ggg/app/http/module/demo"
	"github.com/reallovelei/ggg/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist", "./dist")
	demo.Register(r)
}
