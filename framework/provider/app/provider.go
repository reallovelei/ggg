package app

import (
    "github.com/reallovelei/ggg/framework"
    "github.com/reallovelei/ggg/framework/contract"
)

// GGGAppProvider 提供 App 的具体实现方法
type GGGAppProvider struct {
   BasePath string
}

// Params 获取初始化参数
func (h *GGGAppProvider) Register(container framework.Container) framework.NewInstance {
    return NewApp
}

func (h *GGGAppProvider) Boot(container framework.Container) error {
    return nil
}


func (h *GGGAppProvider) IsDefer() bool {
    return false
}

// Params 获取初始化参数
func (h *GGGAppProvider) Params(container framework.Container) []interface{} {
   return []interface{}{container, h.BasePath}
}

// Name 获取字符串凭证
func (h *GGGAppProvider) Name() string {
    return contract.AppKey
}


