package app

import "github.com/reallovelei/ggg/framework"

// TApp 代表  框架的 App 实现
type HadeApp struct {
    container framework.Container // 服务容器
    baseFolder  string              // 基础路径
}
