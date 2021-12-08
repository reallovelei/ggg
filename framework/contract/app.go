
package contract
// AppKey 定义字符串凭证
const AppKey = "ggg:app"
// App 定义接口
type App interface {
    // Version 定义当前版本
    Version() string
    //BasePath 定义项目基础地址
    BasePath() string

    // ConfigPath 定义了配置文件的路径
    ConfigPath() string
    // LogFolder 定义了日志所在路径
    LogFolder() string
    // ProviderFolder 定义业务自己的服务提供者地址
    ProviderFolder() string
    // MiddlewareFolder 定义业务自己定义的中间件
    MiddlewareFolder() string
    // CommandFolder 定义业务定义的命令
    CommandFolder() string
    // RuntimeFolder 定义业务的运行中间态信息
    RuntimeFolder() string
    // TestFolder 存放测试所需要的信息
    TestFolder() string
}
