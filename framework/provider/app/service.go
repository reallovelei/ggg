package app

import (
	"errors"
	"github.com/google/uuid"
	"github.com/reallovelei/ggg/framework"
	"github.com/reallovelei/ggg/framework/util"
	"path/filepath"
)

// TApp 代表  框架的 App 实现
type GGGApp struct {
	container framework.Container // 服务容器
	basePath  string              // 基础路径
	appId     string              // 表示当前这个app的唯一id, 可以用于分布式锁等
	configMap map[string]string   // 配置
}

// Version 实现版本
func (g GGGApp) Version() string {
	return "0.0.1"
}

// BasePath 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (g GGGApp) BasePath() string {
	if g.basePath != "" {
		return g.basePath
	}

	//// 如果没有设置，从参数获取
	//var basePath string
	//flag.StringVar(&basePath, "base_path", "", "base_path参数，默认为当前路径")
	//flag.Parse()
	//if basePath != "" {
	//	return basePath
	//}

	return util.GetExecDirectory()
}

// ConfigPath  表示配置文件地址
func (g GGGApp) ConfigPath() string {
	// fmt.Println("configmap:", g.configMap)
	if val, ok := g.configMap["config_path"]; ok {
		return val
	}

	return filepath.Join(g.BasePath(), "config")
}

// LogPath 表示日志存放地址
func (g GGGApp) LogPath() string {
	return filepath.Join(g.StoragePath(), "log")
}

func (g GGGApp) HttpPath() string {
	return filepath.Join(g.BasePath(), "app", "web")
}

func (g GGGApp) ConsolePath() string {
	return filepath.Join(g.BasePath(), "app", "command")
}

func (g GGGApp) StoragePath() string {
	return filepath.Join(g.BasePath(), "storage")
}

// ProviderPath 定义业务自己的服务提供者地址
func (g GGGApp) ProviderPath() string {
	if val, ok := g.configMap["provider_path"]; ok {
		return val
	}
	return filepath.Join(g.BasePath(), "app", "provider")
}

// MiddlewarePath 定义业务自己定义的中间件
func (g GGGApp) MiddlewarePath() string {
	return filepath.Join(g.HttpPath(), "middleware")
}

// CommandPath 定义业务定义的命令
func (g GGGApp) CommandPath() string {
	return g.ConsolePath()
}

// RuntimePath 定义业务的运行中间态信息
func (g GGGApp) RuntimePath() string {
	return filepath.Join(g.StoragePath(), "runtime")
}

// TestPath 定义测试需要的信息
func (g GGGApp) TestPath() string {
	return filepath.Join(g.BasePath(), "test")
}

// NewGggApp 初始化 GggApp
func NewApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是 basePath
	container := params[0].(framework.Container)
	basePath := params[1].(string)
	// 如果没有设置，则使用参数
	//if basePath == "" {
	//	flag.StringVar(&basePath, "base_path", "", "base_path参数, 默认为当前路径")
	//	flag.Parse()
	//}

	appId := uuid.New().String()
	configMap := map[string]string{}

	return &GGGApp{basePath: basePath, container: container, appId: appId, configMap: configMap}, nil
}

func (g GGGApp) AppID() string {
	return g.appId
}

func (g GGGApp) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		g.configMap[key] = val
	}
}
