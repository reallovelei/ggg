package kernel

import (
    "github.com/reallovelei/ggg/framework/gin"
    "net/http"
)

type KernelService struct {
    engin *gin.Engine
}

// 初始化 web 引擎服务实例
func NewKernelService(params ...interface{}) (interface{}, error) {
    httpEngine := params[0].(*gin.Engine)
    return &KernelService{engin: httpEngine}, nil
}

// 返回 web 引擎
func (s *KernelService) HttpEngine() http.Handler {
    return s.engin
}
