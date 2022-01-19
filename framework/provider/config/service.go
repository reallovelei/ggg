package config

import (
	"bytes"
	errors2 "errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/reallovelei/ggg/framework"
	"github.com/reallovelei/ggg/framework/contract"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// GGGConfig  表示ggg框架的配置文件服务
type GGGConfig struct {
	c        framework.Container    // 容器
	path     string                 // 文件夹
	keyDelim string                 // 路径的分隔符，默认为点
	lock     sync.RWMutex           // 配置文件读写锁
	envMaps  map[string]string      // 所有的环境变量
	confMaps map[string]interface{} // 配置文件结构，key为文件名
	confRaws map[string][]byte      // 配置文件的原始信息
}

func replaceMap(content []byte, maps map[string]string) []byte {
	if maps == nil {
		return content
	}

	for key, val := range maps {
		reKey := "env(" + key + ")"
		content = bytes.ReplaceAll(content, []byte(reKey), []byte(val))
	}
	return content
}

// 读取某个配置文件
func (conf *GGGConfig) loadConfigFile(path string, file string) error {
	conf.lock.Lock()
	defer conf.lock.Unlock()

	// fmt.Println("loadConfigFile file:", file)

	//  判断文件是否以yaml或者yml作为后缀
	s := strings.Split(file, ".")
	if len(s) == 2 && (s[1] == "yaml" || s[1] == "yml") {
		name := s[0]

		// 读取文件内容
		bf, err := ioutil.ReadFile(filepath.Join(path, file))
		if err != nil {
			return err
		}
		// 直接针对文本做环境变量的替换
		bf = replaceMap(bf, conf.envMaps)

		// 解析对应的文件
		c := map[string]interface{}{}
		if err := yaml.Unmarshal(bf, &c); err != nil {
			return err
		}

		conf.confMaps[name] = c
		conf.confRaws[name] = bf

		// fmt.Println("loadConfigFile 2", conf.confMaps)
		// 读取app.path中的信息，更新app对应的folder
		if name == "app" && conf.c.IsBind(contract.AppKey) {
			if p, ok := c["path"]; ok {
				appService := conf.c.MustMake(contract.AppKey).(contract.App)
				appService.LoadAppConfig(cast.ToStringMapString(p))
			}
		}
	}
	return nil
}

// 删除文件的操作
func (conf *GGGConfig) removeConfigFile(folder string, file string) error {
	conf.lock.Lock()
	defer conf.lock.Unlock()
	s := strings.Split(file, ".")
	// 只有yaml或者yml后缀才执行
	if len(s) == 2 && (s[1] == "yaml" || s[1] == "yml") {
		name := s[0]
		// 删除内存中对应的key
		delete(conf.confRaws, name)
		delete(conf.confMaps, name)
	}
	return nil
}

func NewGGGConfig(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	envPath := params[1].(string)
	envMaps := params[2].(map[string]string)

	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		return nil, errors2.New("folder " + envPath + "not exist: " + err.Error())
	}

	// 实例化
	conf := &GGGConfig{
		c:        container,
		path:     envPath,
		keyDelim: ".",
		lock:     sync.RWMutex{},
		envMaps:  envMaps,
		confMaps: map[string]interface{}{},
		confRaws: map[string][]byte{},
	}
	// 读取目录的每个文件
	files, err := ioutil.ReadDir(envPath)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 挨个加载配置文件
	for _, file := range files {
		fileName := file.Name()
		err := conf.loadConfigFile(envPath, fileName)

		if err != nil {
			log.Println(err)
			continue
		}
	}

	// 监控文件夹文件
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	// 添加监听目录
	err = watch.Add(envPath)
	if err != nil {
		return nil, err
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		for {
			select {
			case event := <-watch.Events:
				{
					// 判断事件类型
					path, _ := filepath.Abs(event.Name)
					index := strings.LastIndex(path, string(os.PathSeparator))
					folder := path[:index]
					fileName := path[index+1:]

					// 创建文件
					if event.Op&fsnotify.Create == fsnotify.Create {
						log.Println("创建文件:", event.Name)
						conf.loadConfigFile(folder, fileName)
					}

					// 写入文件
					if event.Op&fsnotify.Write == fsnotify.Write {
						log.Println("写入文件:", event.Name)
						conf.loadConfigFile(folder, fileName)
					}

					// 删除文件
					if event.Op&fsnotify.Remove == fsnotify.Remove {
						log.Println("创建文件:", event.Name)
						conf.removeConfigFile(folder, fileName)
					}
				}
			case err := <-watch.Errors:
				{
					log.Println("error :", err)
					return
				}
			}
		}
	}()

	return conf, nil
}

func searchMap(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}
	// 判断是否有下个路径
	next, ok := source[path[0]]

	if ok {
		if len(path) == 1 {
			return next
		}

		switch next.(type) {
		case map[interface{}]interface{}:
			// 如果是 interface的map, 使用cast 进行下 value 的转换
			return searchMap(cast.ToStringMap(next), path[1:])

		case map[string]interface{}:
			// 如果是map[string], 直接循环调用
			return searchMap(next.(map[string]interface{}), path[1:])

		default:
			// 否则 return nil
			return nil
		}
	}
	return nil
}

// 通过path来获取某个配置项
func (conf *GGGConfig) find(key string) interface{} {
	conf.lock.RLock()
	defer conf.lock.RUnlock()
	// fmt.Println("find", conf.envMaps)
	return searchMap(conf.confMaps, strings.Split(key, conf.keyDelim))
}

// IsExist check setting is exist
func (conf *GGGConfig) IsExist(key string) bool {
	return conf.find(key) != nil
}

// Get 获取某个配置项
func (conf *GGGConfig) Get(key string) interface{} {
	return conf.find(key)
}

// GetBool 获取bool类型配置
func (conf *GGGConfig) GetBool(key string) bool {
	return cast.ToBool(conf.find(key))
}

// GetInt 获取int类型配置
func (conf *GGGConfig) GetInt(key string) int {
	return cast.ToInt(conf.find(key))
}

// GetFloat64 get float64
func (conf *GGGConfig) GetFloat64(key string) float64 {
	return cast.ToFloat64(conf.find(key))
}

// GetTime get time type
func (conf *GGGConfig) GetTime(key string) time.Time {
	return cast.ToTime(conf.find(key))
}

// GetString get string typen
func (conf *GGGConfig) GetString(key string) string {
	// fmt.Println("find", key, conf.find(key))
	return cast.ToString(conf.find(key))
}

// GetIntSlice get int slice type
func (conf *GGGConfig) GetIntSlice(key string) []int {
	return cast.ToIntSlice(conf.find(key))
}

// GetStringSlice get string slice type
func (conf *GGGConfig) GetStringSlice(key string) []string {
	return cast.ToStringSlice(conf.find(key))
}

// GetStringMap get map which key is string, value is interface
func (conf *GGGConfig) GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(conf.find(key))
}

// GetStringMapString get map which key is string, value is string
func (conf *GGGConfig) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(conf.find(key))
}

// GetStringMapStringSlice get map which key is string, value is string slice
func (conf *GGGConfig) GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(conf.find(key))
}

// Load a config to a struct, val should be an pointer
func (conf *GGGConfig) Load(key string, val interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "yaml",
		Result:  val,
	})
	if err != nil {
		return err
	}

	return decoder.Decode(conf.find(key))
}
