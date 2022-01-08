package framework

import (
	"errors"
	"fmt"
	"sync"
)

// Container 是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回 error
	Bind(provider ServiceProvider) error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务，
	Make(key string) (interface{}, error)
	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会 panic。
	// 所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}
	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的 params 参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// GGGContainer 是服务容器的具体实现
type GGGContainer struct {
	Container // 强制要求 GGGContainer 实现 Container 接口
	// providers 存储注册的服务提供者，key 为字符串凭证 bind 提供者
	providers map[string]ServiceProvider
	// instance 存储具体的服务实例，key 为字符串凭证
	instances map[string]interface{}
	// lock 用于锁住对容器的变更操作, Bind 是一次性的，但是 Make 是频繁的。所以使用读写锁的性能会优于互斥锁。
	lock sync.RWMutex
}

// NewContainer 创建一个服务容器
func NewContainer() *GGGContainer {
	return &GGGContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 输出服务容器中注册的关键字
func (ggg *GGGContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range ggg.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

// Bind 将服务容器和关键字做了绑定
func (ggg *GGGContainer) Bind(provider ServiceProvider) error {
	// Bind 是写操作, 所以一开头先加上 一把写锁。 RLock是读锁。
	ggg.lock.Lock()

	key := provider.Name()
	// 修改 providers 这个字段，它的 key 为关键字，value 为注册的 ServiceProvider。
	ggg.providers[key] = provider

	ggg.lock.Unlock()

	// fmt.Println("privider Name:", key, ggg.providers)
	// if provider is not defer
	// 如果这个服务实例要延迟实例化，即等到第一次 make 的时候再实例化，那么在 Bind 操作的时候，就什么都不需要做。
	// 如果不是延迟实例化 就需要做如下事情
	if provider.IsDefer() == false {
		// ServiceProvider 定义过一个 Boot 方法，是为了服务实例化前做一些准备工作的。所以在实例化之前，要先调用这个 Boot 方法。
		if err := provider.Boot(ggg); err != nil {
			return err
		}

		// 实例化方法
		params := provider.Params(ggg)

		method := provider.Register(ggg)

		// fmt.Println("privider 3 Name:", key)
		instance, err := method(params...)
		if err != nil {
			fmt.Println("privider ERROR:", err.Error())
			return errors.New(err.Error())
		}

		fmt.Println("privider 4 Name:", key)
		ggg.instances[key] = instance
	}

	return nil
}

func (ggg *GGGContainer) IsBind(key string) bool {
	return ggg.findServiceProvider(key) != nil
}

func (ggg *GGGContainer) findServiceProvider(key string) ServiceProvider {
	ggg.lock.RLock()
	defer ggg.lock.RUnlock()
	if sp, ok := ggg.providers[key]; ok {
		return sp
	}
	return nil
}

func (ggg *GGGContainer) Make(key string) (interface{}, error) {
	return ggg.make(key, nil, false)
}

func (ggg *GGGContainer) MustMake(key string) interface{} {
	serv, err := ggg.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (ggg *GGGContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ggg.make(key, params, true)
}

func (ggg *GGGContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// force new a
	if err := sp.Boot(ggg); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(ggg)
	}
	method := sp.Register(ggg)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

// 真正的实例化一个服务
func (ggg *GGGContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	ggg.lock.RLock()
	defer ggg.lock.RUnlock()
	// 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	sp := ggg.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return ggg.newInstance(sp, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := ggg.instances[key]; ok {
		return ins, nil
	}

	// 容器中还未实例化，则进行一次实例化
	inst, err := ggg.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	ggg.instances[key] = inst
	return inst, nil
}
