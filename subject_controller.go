package main

import (
	"fmt"
	"github.com/reallovelei/ggg/app/provider/demo"
	"github.com/reallovelei/ggg/framework/gin"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")
}

// 对应路由 /subject/list/all
func SubjectListController(c *gin.Context) {
	// 获取demo服务实例
	demoService := c.MustMake(demo.DemoKey).(demo.Service)

	// 调用服务实例的方法
	foo := demoService.GetAllStudent()

	// 输出结果
	c.ISetOkStatus().IJson(foo)
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	subjectId, _ := c.DefaultParamInt("id", 0)
	c.ISetOkStatus().IJson("ok, SubjectGetController:" + fmt.Sprint(subjectId))

}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}
