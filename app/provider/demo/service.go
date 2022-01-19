package demo

import (
	"fmt"
	"github.com/reallovelei/ggg/framework"
)

const (
	_selectByUpdate = "SELECT push_id, updated_at FROM push_logs where updated_at < ?"
	_deleteByUpdate = "DELETE FROM push_logs where updated_at < ?"
)

// 具体的接口实例
type DemoService struct {
	// 实现接口
	Service

	// 参数
	c framework.Container
}

// 初始化实例的方法
func NewDemoService(params ...interface{}) (interface{}, error) {
	// 这里需要将参数展开
	c := params[0].(framework.Container)

	fmt.Println("new demo service")
	// 返回实例
	return &DemoService{c: c}, nil
}

// 实现接口
func (s *DemoService) GetStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "zhanglei",
		},
		{
			ID:   2,
			Name: "chenjiang",
		},
	}
}
