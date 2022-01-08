package command

import (
	"fmt"
	"github.com/kr/pretty"
	"github.com/reallovelei/ggg/framework/cobra"
	"github.com/reallovelei/ggg/framework/contract"
)

// 一个Command 可以add 另一个Command
func initConfigCommand() *cobra.Command {
	configCommand.AddCommand(configGetCommand)
	return configCommand
}

// envCommand 获取当前的App环境
var configCommand = &cobra.Command{
	Use:   "config",
	Short: "获取配置相关信息",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

// evnListCommand 获取所有App环境变量
var configGetCommand = &cobra.Command{
	Use:   "get",
	Short: "获取某个配置信息",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		configService := container.MustMake(contract.ConfigKey).(contract.Config)

		if len(args) != 1 {
			fmt.Println("参数错误")
			return nil
		}

		configPath := args[0]
		value := configService.Get(configPath)
		if value == nil {
			fmt.Println("配置路径", configPath, "不存在")
			return nil
		}

		fmt.Printf("%# v\n", pretty.Formatter(value))
		return nil
	},
}
