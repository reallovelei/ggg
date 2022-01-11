package main

import (
	"github.com/reallovelei/ggg/app/command"
	"github.com/reallovelei/ggg/app/web"
	"github.com/reallovelei/ggg/framework"
	"github.com/reallovelei/ggg/framework/provider/app"
	"github.com/reallovelei/ggg/framework/provider/config"
	distributed "github.com/reallovelei/ggg/framework/provider/distribute"
	"github.com/reallovelei/ggg/framework/provider/env"
	"github.com/reallovelei/ggg/framework/provider/kernel"
	"github.com/reallovelei/ggg/framework/provider/log"
	"github.com/reallovelei/ggg/framework/provider/orm"
)

func main() {
	// 初始化服务容器
	container := framework.NewContainer()

	// 绑定服务提供者
	container.Bind(&app.GGGAppProvider{})

	container.Bind(&env.GGGEnvProvider{})

	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.GGGConfigProvider{})

	container.Bind(&log.GGGLogServiceProvider{})

	container.Bind(&orm.GormProvider{})
	if engine, err := web.NewHttpEngine(); err == nil {
		container.Bind(&kernel.GGGKernelProvider{HttpEngine: engine})
	}

	command.RunCommand(container)

	// 创建 engine 结构
	//core := gin.New()
	//
	//
	//// bind service
	//core.Bind(&app.GGGAppProvider{BasePath:"/tmp"})
	//
	//core.Use(gin.Recovery())
	//
	//core.Use(middleware.Cost())
	//RegisterRouter(core)
	//
	//server := &web.Server {
	//    Handler:core,
	//    Addr:":8888",
	//}
	//
	//// 使用一个协程来 监听服务
	//go func() {
	//    server.ListenAndServe()
	//}()
	//
	//quit := make(chan os.Signal)
	//// 监控信号：SIGINT, SIGTERM, SIGQUIT
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//// 这里会阻塞当前goroutine等待信号
	//<-quit
	//
	//// 调用Server.Shutdown graceful结束
	//timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//if err := server.Shutdown(timeoutCtx); err != nil {
	//    log.Fatal("Server Shutdown:", err)
	//}
}
