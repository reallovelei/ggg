package env

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/reallovelei/ggg/framework/contract"
	"io"
	"os"
	"path"
	"strings"
)

type GGGEnv struct {
	path string            // .env 的目录
	maps map[string]string // 保存环境变量
}

//gggEnv := &GGGEnv{
//    path: path,
//    // 实例化环境变量，APP_ENV 默认设置为开发环境
//    maps: map[string]string{"APP_ENV": contract.EnvDev},
//}

// NewEnv 有一个参数，.env文件所在的目录
// example: NewEnv("/envfolder/") 会读取文件: /envfolder/.env
// .env的文件格式 FOO_ENV=BAR
func NewEnv(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("NewEnv param error")
	}

	fmt.Println("will new env")
	// 读取folder文件
	envPath := params[0].(string)
	fmt.Println("will new env path", envPath)
	// 实例化
	env := &GGGEnv{
		path: envPath,
		// 实例化环境变量，APP_ENV默认设置为开发环境
		maps: map[string]string{"APP_ENV": contract.EnvDev},
	}

	// 解析folder/.env文件
	file := path.Join(envPath, ".env")
	// 读取.env文件, 不管任意失败，都不影响后续

	// 打开文件.env
	fi, err := os.Open(file)
	if err == nil {
		defer fi.Close()

		// 读取文件
		br := bufio.NewReader(fi)
		for {
			// 按照行进行读取
			line, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			// 按照等号解析
			s := bytes.SplitN(line, []byte{'='}, 2)
			// 如果不符合规范，则过滤
			if len(s) < 2 {
				continue
			}
			// 保存map
			key := string(s[0])
			val := string(s[1])
			env.maps[key] = val
		}
	}

	// 获取当前程序的环境变量，并且覆盖.env文件下的变量
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) < 2 {
			continue
		}
		env.maps[pair[0]] = pair[1]
	}

	// 返回实例
	return env, nil
}

// 返回当前的环境变量，每个env 配置文件都要有
func (env *GGGEnv) AppEnv() string {
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
