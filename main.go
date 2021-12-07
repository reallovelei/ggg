package main

import (
    "context"
    "log"
    "time"

    // "context"
    "github.com/gohade/hade/framework/gin"
    "github.com/gohade/hade/framework/middleware"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    core := gin.New()
    core.Use(gin.Recovery())
    core.Use(middleware.Cost())

    registerRouter(core)

    server := &http.Server {
        Handler:core,
        Addr:":9595",
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
