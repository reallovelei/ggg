package app

import (
	"errors"
	"flag"
	"github.com/reallovelei/ggg/framework"
	"github.com/reallovelei/ggg/framework/util"
	"path/filepath"
)

// TApp 代表  框架的 App 实现
type GGGApp struct {
	container framework.Container // 服务容器
	basePath  string              // 基础路径
}

// Version 实现版本
func (h GGGApp) Version() string {
	return "0.0.1"
}

// BasePath 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (h GGGApp) BasePath() string {
	if h.basePath != "" {
		return h.basePath
	}

	// 如果没有设置，从参数获取
	var basePath string
	flag.StringVar(&basePath, "base_path", "", "base_path参数，默认为当前路径")
	flag.Parse()
	if basePath != "" {
		return basePath
	}

	return util.GetExecDirectory()
}

// ConfigPath  表示配置文件地址
func (h GGGApp) ConfigPath() string {
	return filepath.Join(h.BasePath(), "config")
}

// LogPath 表示日志存放地址
func (h GGGApp) LogPath() string {
	return filepath.Join(h.StoragePath(), "log")
}

func (h GGGApp) HttpPath() string {
	return filepath.Join(h.BasePath(), "web")
}

func (h GGGApp) ConsolePath() string {
	return filepath.Join(h.BasePath(), "console")
}

func (h GGGApp) StoragePath() string {
	return filepath.Join(h.BasePath(), "storage")
}

// ProviderPath 定义业务自己的服务提供者地址
func (h GGGApp) ProviderPath() string {
	return filepath.Join(h.BasePath(), "provider")
}

// MiddlewarePath 定义业务自己定义的中间件
func (h GGGApp) MiddlewarePath() string {
	return filepath.Join(h.HttpPath(), "middleware")
}

// CommandPath 定义业务定义的命令
func (h GGGApp) CommandPath() string {
	return filepath.Join(h.ConsolePath(), "command")
}

// RuntimePath 定义业务的运行中间态信息
func (h GGGApp) RuntimePath() string {
	return filepath.Join(h.StoragePath(), "runtime")
}

// TestPath 定义测试需要的信息
func (h GGGApp) TestPath() string {
	return filepath.Join(h.BasePath(), "test")
}

// NewGggApp 初始化 GggApp
func NewApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}
	// 有两个参数，一个是容器，一个是 basePath
	container := params[0].(framework.Container)
	basePath := params[1].(string)
	return &GGGApp{basePath: basePath, container: container}, nil
}
