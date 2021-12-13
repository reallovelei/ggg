package demo

import (
	"fmt"
	"github.com/reallovelei/ggg/framework"
)

type DemoServiceProvider struct {
}

// 获得服务凭证
func (sp *DemoServiceProvider) Name() string {
	return DemoKey
}

//
func (sp *DemoServiceProvider) Register(c framework.Container) framework.NewInstance {
	return NewDemoService
}

func (sp *DemoServiceProvider) IsDefer() bool {
	return true
}

func (sp *DemoServiceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *DemoServiceProvider) Boot(c framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}
