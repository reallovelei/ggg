package env

import "errors"

type GGGEnv struct {
    path string             // .env 的目录
    maps map[string]string  // 保存环境变量
}

gggEnv := &GGGEnv{
    path: path,
    // 实例化环境变量，APP_ENV 默认设置为开发环境
    maps: map[string]string{"APP_ENV": contract.EnvDevelopment},
}

func NewEnv(params ...interface{}) (interface{}, error) {
    if len(params) != 1 {
        return nil, errors.New("NewEnv param error")
    }


}

// 返回当前的环境变量，每个env 配置文件都要有
func (env *GGGEnv) AppEnv(key string) string {
    return env.Get("APP_ENV")
}

func (env *GGGEnv) IsExist(key string) bool {
    _, ok := env.maps[key]
    return ok
}

func (env *GGGEnv) Get(key string) string {
    if val, ok := env.maps[key]; ok {
        return val
    }
    return ""
}

func (env *GGGEnv) All() map[string]string {
    return env.maps
}


