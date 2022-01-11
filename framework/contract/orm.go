package contract

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/reallovelei/ggg/framework"
	"gorm.io/gorm"
	"net"
	"strconv"
	"time"
)

// 代表orm的服务
const ORMKey = "ggg:orm"

// ORMService 表示传入的参数
type ORMService interface {
	GetDB(option ...DBOption) (*gorm.DB, error)
}

type DBOption func(container framework.Container, config *DBConfig) error

// DBConfig 代表数据库连接的所有配置
type DBConfig struct {
	// 以下配置关于dsn
	WriteTimeout         string `yaml:"write_timeout"` // 写超时时间
	Loc                  string `yaml:"loc"`           // 时区
	Port                 int    `yaml:"port"`          // 端口
	ReadTimeout          string `yaml:"read_timeout"`  // 读超时时间
	Charset              string `yaml:"charset"`       // 字符集
	ParseTime            bool   `yaml:"parse_time"`    // 是否解析时间
	Protocol             string `yaml:"protocol"`      // 传输协议
	Dsn                  string `yaml:"dsn"`           // 直接传递dsn，如果传递了，其他关于dsn的配置均无效
	Database             string `yaml:"database"`      // 数据库
	Collation            string `yaml:"collation"`     // 字符序
	Timeout              string `yaml:"timeout"`       // 连接超时时间
	Username             string `yaml:"username"`      // 用户名
	Password             string `yaml:"password"`      // 密码
	Driver               string `yaml:"driver"`        // 驱动
	Host                 string `yaml:"host"`          // 数据库地址
	AllowNativePasswords bool   `yaml:"allow_native_passwords"`

	// 以下配置关于连接池
	ConnMaxIdle     int    `yaml:"conn_max_idle"`     // 最大空闲连接数
	ConnMaxOpen     int    `yaml:"conn_max_open"`     // 最大连接数
	ConnMaxLifetime string `yaml:"conn_max_lifetime"` // 连接最大生命周期
	ConnMaxIdletime string `yaml:"conn_max_idletime"` // 空闲最大生命周期

	// 以下配置关于gorm   继承了 *gorm.Config 的所有方法
	*gorm.Config // 集成gorm的配置
}

// FormatDsn 生成dsn
func (conf *DBConfig) FormatDsn() (string, error) {
	fmt.Printf("formatDsn start!~ conf: %+v \n", conf)

	port := strconv.Itoa(conf.Port)

	timeout, err := time.ParseDuration(conf.Timeout)
	if err != nil {
		fmt.Println("conf 1", err.Error())
		return "", err
	}
	readTimeout, err := time.ParseDuration(conf.ReadTimeout)
	if err != nil {
		fmt.Println("conf 2", err.Error())
		return "", err
	}
	writeTimeout, err := time.ParseDuration(conf.WriteTimeout)
	if err != nil {
		fmt.Println("conf 3", err.Error())
		return "", err
	}
	location, err := time.LoadLocation(conf.Loc)
	if err != nil {
		fmt.Println("conf 4", err.Error())
		return "", err
	}
	driverConf := &mysql.Config{
		User:                 conf.Username,
		Passwd:               conf.Password,
		Net:                  conf.Protocol,
		Addr:                 net.JoinHostPort(conf.Host, port),
		DBName:               conf.Database,
		Collation:            conf.Collation,
		Loc:                  location,
		Timeout:              timeout,
		ReadTimeout:          readTimeout,
		WriteTimeout:         writeTimeout,
		ParseTime:            conf.ParseTime,
		AllowNativePasswords: conf.AllowNativePasswords,
	}
	fmt.Printf("conf:%+v \n", driverConf)

	return driverConf.FormatDSN(), nil
}
