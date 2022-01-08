package env

import (
	"github.com/reallovelei/ggg/framework"
	"github.com/reallovelei/ggg/framework/contract"
)

type GGGEnvProvider struct {
	Path string
}

func (provider *GGGEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewGGGEnv
}

func (provider *GGGEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Path = app.BasePath()
	return nil
}

func (provider *GGGEnvProvider) IsDefer() bool {
	return false
}

func (provider *GGGEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Path}
}

func (provider *GGGEnvProvider) Name() string {
	return contract.EnvKey
}
