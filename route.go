package main

import (
    "github.com/gohade/hade/framework/gin"
    "github.com/gohade/hade/framework/middleware"
)
// 注册路由规则
func registerRouter(core *gin.Engine) {
    core.GET("/user/login", middleware.Test3(), UserLoginController)
}
