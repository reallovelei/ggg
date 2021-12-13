package service

import (
	"github.com/reallovelei/ggg/framework"
	"github.com/reallovelei/ggg/framework/contract"
)

// GGGLog 的通用实例
type GGGLog struct {
	// 五个必要参数
	level      contract.LogLevel   // 日志级别
	formatter  contract.Formatter  // 日志格式化方法
	ctxFielder contract.CtxFielder // ctx获取上下文字段
	output     io.Writer           // 输出
	c          framework.Container // 容器
}
