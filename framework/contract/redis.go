
package contract

import (
    "fmt"
    "github.com/go-redis/redis/v8"
    "github.com/reallovelei/ggg/framework"
)


const RedisKey = "ggg:redis"

// RedisConfig 为 ggg 定义的Redis配置结构
type RedisConfig struct {
    *redis.Options
}

// RedisOption 代表初始化的时候的选项
type RedisOption func(container framework.Container, config *RedisConfig) error

// RedisService 表示一个redis服务
type RedisService interface {
    // GetClient 获取redis连接实例
    // 可变参数 RedisOption，这个可变参数是一个函数结构.
    // 参数中带有传递进入了的 RedisConfig 指针
    GetClient(option ...RedisOption) (*redis.Client, error)
}

// UniqKey 用来唯一标识一个RedisConfig配置
func (config *RedisConfig) UniqKey() string {
    return fmt.Sprintf("%v_%v_%v_%v", config.Addr, config.DB, config.Username, config.Network)
}
