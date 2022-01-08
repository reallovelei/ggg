package config

import (
	"github.com/reallovelei/ggg/framework"
	"github.com/reallovelei/ggg/framework/contract"
	"path/filepath"
)

type GGGConfigProvider struct {
}

// Register registe a new function for make a service instance
func (provider *GGGConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewGGGConfig
}

// Boot will called when the service instantiate
func (provider *GGGConfigProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *GGGConfigProvider) IsDefer() bool {
	return false
}

// Name define the name for this service
func (provider *GGGConfigProvider) Name() string {
	return contract.ConfigKey
}

func (provider *GGGConfigProvider) Params(c framework.Container) []interface{} {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()

	configPath := appService.ConfigPath()
	envPath := filepath.Join(configPath, env)
	return []interface{}{c, envPath, envService.All()}
}
