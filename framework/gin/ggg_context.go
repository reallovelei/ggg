package gin

import "github.com/gohade/hade/framework"

//  实现在engine 对  container 的绑定封装
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
    return engine.container.Bind(provider)
}

// 实现在engine中  IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
    return engine.container.IsBind(key)
}


// context 实现 container 的几个封装
// 实现 make 的封装
func (ctx *Context) Make(key string) (interface{}, error) {
    return ctx.container.Make(key)
}

// 实现 mustMake 的封装
func (ctx *Context) MustMake(key string) interface{} {
    return ctx.container.MustMake(key)
}

// 实现 makenew 的封装
func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
    return ctx.container.MakeNew(key, params)
}

