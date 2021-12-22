package env

import "github.com/reallovelei/ggg/framework"

type EngProvider struct {
	Path string
}

func (provider *EngProvider) Register(c framework.Container) framework.NewInstance {
	return New
}
