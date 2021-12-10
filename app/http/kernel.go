package http

import "github.com/reallovelei/ggg/framework/gin"

func NewHttpEngin() (*gin.Engine, error) {
	// 设置为Release模式，默认在启动中不输出调试信息
	gin.SetMode(gin.ReleaseMode)
	// 默认启动一个web 引擎
	r := gin.Default()
	// 业务绑定路由操作

	Routes(r)
	// 返回绑定路由后的web 引擎
	return r, nil
}
