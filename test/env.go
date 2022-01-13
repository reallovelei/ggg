package tests

import (
	"github.com/reallovelei/ggg/framework"
	"github.com/reallovelei/ggg/framework/provider/app"
	//	"github.com/reallovelei/ggg/framework/provider/env"
)

const (
	// todo config ?
	BasePath = "/Users/Ben/work/wangxiao/ggg-frame/framework"
)

func InitBaseContainer() framework.Container {
	// 初始化服务容器
	container := framework.NewContainer()
	// 绑定App服务提供者
	container.Bind(&app.GGGAppProvider{BasePath: BasePath})
	// 后续初始化需要绑定的服务提供者...
	//	container.Bind(&env.GGGTestingEnvProvider{})
	return container
}
