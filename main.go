package main

import (
    "context"
    "github.com/reallovelei/ggg/app/provider/demo"
    "github.com/reallovelei/ggg/framework/middleware"
    "log"
    "time"

    // "context"
    "github.com/reallovelei/ggg/framework/gin"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // 创建 engine 结构
    core := gin.New()


    // bind service
    core.Bind(&demo.DemoServiceProvider{})

    core.Use(gin.Recovery())

    core.Use(middleware.Cost())
    RegisterRouter(core)

    server := &http.Server {
        Handler:core,
        Addr:":8888",
    }

    // 使用一个协程来 监听服务
    go func() {
        server.ListenAndServe()
    }()

    quit := make(chan os.Signal)
    // 监控信号：SIGINT, SIGTERM, SIGQUIT
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    // 这里会阻塞当前goroutine等待信号
    <-quit

    // 调用Server.Shutdown graceful结束
    timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(timeoutCtx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }
}
