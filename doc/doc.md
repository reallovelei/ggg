## 规范
服务的接口协议，统一放在框架目录下的 framework/contract。
对应的服务提供者和服务实现，统一存放在框架目录下的 framework/provider 中。

将所有框架级别的接口协议放在 framework/contract 中的设计有两个好处:
* 框架协议的关键字，我希望使用 contract.xxx 这个语义来区分，比如 App 服务的接口为 contract.App、日志服务的接口为 contract.Log，它们的 namespace 都是 contract，这样在使用的时候记忆成本会比较低。
* 将框架提供的所有接口协议都放在一个文件夹中，在阅读框架提供哪些服务的时候，也更清晰明了。


## 日志
归并为下列七种日志级别：
* panic，表示会导致整个程序出现崩溃的日志信息。
* fatal，表示会导致当前这个请求出现提前终止的错误信息。
* error，表示出现错误，但是不一定影响后续请求逻辑的错误信息。
* warn，表示出现错误，但是一定不影响后续请求逻辑的报警信息。
* info，表示正常的日志信息输出。
* debug，表示在调试状态下打印出来的日志信息。
* trace，表示最详细的信息，一般信息量比较大，可能包含调用堆栈等信息在。

error 级别之上，我们把导致程序崩溃和导致请求结束的错误拆分出来，分为 panic 和 fatal 两个类型来定义级别。
而其他的 error、warn、info、debug 都和其他的日志系统一致。
另外也增加一个 trace 级别，当需要打印调用堆栈等这些比较详细的信息的时候，可以使用这种日志级别。

## 开发
我们将服务容器 container 存放在根 Command 中，让所有的命令都可以通过 Root() 方法获取到根 Command，再获取到 container。

##
```
// 编译
编译前端  ./ggg build frontend
编译后端  ./ggg build backend
同时编译前后端 ./ggg build all

// 运行 http
./ggg app start

// 定时任务相关
./ggg cron
```


