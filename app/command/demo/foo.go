package demo

import (
    "fmt"
    "github.com/reallovelei/ggg/framework/cobra"
    "github.com/reallovelei/ggg/framework/contract"
    "github.com/reallovelei/ggg/framework/provider/log/service"
    "log"
)

// InitFoo 初始化Foo命令

func InitFoo() *cobra.Command {
	FooCommand.AddCommand(Foo1Command)
	return FooCommand
}

// FooCommand 代表Foo命令
var FooCommand = &cobra.Command{
	Use:     "foo",
	Short:   "foo的简要说明",
	Long:    "foo的长说明",
	Aliases: []string{"fo", "f"},
	Example: "foo命令的例子",
	RunE: func(c *cobra.Command, args []string) error {
		configService := c.GetContainer().MustMake(contract.ConfigKey).(contract.Config)
		envService := c.GetContainer().MustMake(contract.EnvKey).(contract.Env)

        logger := c.GetContainer().MustMake(contract.LogKey).(*service.GGGConsoleLog)
		fmt.Println("APP_ENV: ", envService.Get("APP_ENV"))
		fmt.Println("FOO_ENV: ", envService.Get("FOO_ENV"))
		fmt.Println("config url:", configService.GetString("app.url"))

		nums := []int{1, 2, 3, 4, 5, 6}
		fmt.Println("     len cap   address")
		fmt.Print("111---", len(nums), cap(nums))
		fmt.Printf("    %p\n", nums) //0xc4200181e0


		logger.Info(c.Context(),"This is Demo Command", nil)

		//log.Println(container)
		return nil
	},
}

// Foo1Command 代表Foo命令的子命令Foo1
var Foo1Command = &cobra.Command{
	Use:     "foo1",
	Short:   "foo1的简要说明",
	Long:    "foo1的长说明",
	Aliases: []string{"fo1", "f1"},
	Example: "foo1命令的例子",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()

        appService := container.MustMake(contract.AppKey).(contract.App)
        appService.LogPath()

		log.Println("This is Demo1 Command  LogPath:"+appService.LogPath(), container)
		return nil
	},
}
