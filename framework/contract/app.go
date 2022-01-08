package contract

// AppKey 定义字符串凭证
const AppKey = "ggg:app"

// App 定义接口
type App interface {
	AppID() string
	// Version 定义当前版本
	Version() string
	//BasePath 定义项目基础地址
	BasePath() string

	// ConfigPath 定义了配置文件的路径
	ConfigPath() string
	// LogPath 定义了日志所在路径
	LogPath() string
	// ProviderPath 定义业务自己的服务提供者地址
	ProviderPath() string
	// MiddlewarePath 定义业务自己定义的中间件
	MiddlewarePath() string
	// CommandPath 定义业务定义的命令
	CommandPath() string
	// RuntimePath 定义业务的运行中间态信息
	RuntimePath() string
	// TestPath 存放测试所需要的信息
	TestPath() string
	// LoadAppConfig 加载新的AppConfig，key为对应的函数转为小写下划线，比如ConfigFolder => config_folder
	LoadAppConfig(kv map[string]string)
}
