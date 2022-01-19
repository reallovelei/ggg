package demo

// 合约文件主要做两件事情
// 1. 定义服务的关键字凭证，用这个凭证来注册到容器中。 这里定义为 "ggg:demo"
// 2. 定义接口

// Demo 服务的 key  就是所谓的凭证。
const DemoKey = "ggg:demo"

// Demo 服务的接口
type Service interface {
	GetAllStudent() []Student
	DelPushLog()
}

// Demo 服务接口定义的一个数据结构
type Student struct {
	ID   int
	Name string
}
