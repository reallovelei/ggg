package gin

import "github.com/reallovelei/ggg/framework"

func (engine *Engine) SetContainer(container framework.Container) {
	engine.container = container
}

//func (engine *Engine) Bind(provider framework.ServiceProvider) error {
//    return engine.container.Bind(provider)
//}
//
//func (engine *Engine) IsBind(key string) bool {
//    return engine.IsBind(key)
//}

//  实现在engine 对  container 的绑定封装
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

// 实现在engine中  IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}
