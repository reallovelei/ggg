package kernel

import (
	"github.com/reallovelei/ggg/framework"
	"github.com/reallovelei/ggg/framework/contract"
	"github.com/reallovelei/ggg/framework/gin"
)

// GGGKernelProvider 提供web引擎
type GGGKernelProvider struct {
	HttpEngine *gin.Engine
}

// Register 注册服务提供者
func (provider *GGGKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewKernelService
}

// Boot 启动的时候判断是否由外界注入了Engine，如果注入的化，用注入的，如果没有，重新实例化
func (provider *GGGKernelProvider) Boot(c framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
}

// IsDefer 引擎的初始化我们希望开始就进行初始化
func (provider *GGGKernelProvider) IsDefer() bool {
	return false
}

// Params 参数就是一个HttpEngine
func (provider *GGGKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}

// Name 提供凭证
func (provider *GGGKernelProvider) Name() string {
	return contract.KernelKey
}
