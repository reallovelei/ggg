package main

import (
    "github.com/reallovelei/ggg/framework/gin"
    "github.com/reallovelei/ggg/framework/middleware"
)
// 注册路由规则
func RegisterRouter(core *gin.Engine) {
    // core.GET("/user/login", middleware.Test3(), UserLoginController)

    // 批量通用前缀
    subjectApi := core.Group("/subject")
    {
        subjectApi.Use(middleware.Test3())

        // 动态路由
        // subjectApi.DELETE("/:id", SubjectDelController)
        // subjectApi.PUT("/:id", SubjectUpdateController)
        // subjectApi.GET("/:id", middleware.Test3(), SubjectGetController)
        subjectApi.GET("/list/all", SubjectListController)

        //subjectInnerApi := subjectApi.Group("/info")
        //{
        //    subjectInnerApi.GET("/name", SubjectNameController)
        //}
    }
}
