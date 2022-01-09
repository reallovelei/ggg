## 规范
服务的接口协议，统一放在框架目录下的 framework/contract。
对应的服务提供者和服务实现，统一存放在框架目录下的 framework/provider 中。

将所有框架级别的接口协议放在 framework/contract 中的设计有两个好处:
* 框架协议的关键字，我希望使用 contract.xxx 这个语义来区分，比如 App 服务的接口为 contract.App、日志服务的接口为 contract.Log，它们的 namespace 都是 contract，这样在使用的时候记忆成本会比较低。
* 将框架提供的所有接口协议都放在一个文件夹中，在阅读框架提供哪些服务的时候，也更清晰明了。

## 开发
我们将服务容器 container 存放在根 Command 中，让所有的命令都可以通过 Root() 方法获取到根 Command，再获取到 container。

##
```
// 编译
go build
// 运行 http
./ggg app start

// 定时任务相关
./ggg cron
```
