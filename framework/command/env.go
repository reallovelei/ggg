package command

import (
	"fmt"
	"github.com/reallovelei/ggg/framework/cobra"
	"github.com/reallovelei/ggg/framework/contract"
	"github.com/reallovelei/ggg/framework/util"
)

func initEnvCommand() *cobra.Command {
	envCommand.AddCommand(envListCommand)
	return envCommand
}

var envCommand = &cobra.Command{
	Short: "获取当前的App环境",
	Use:   "env",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		// 获取env 环境
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		// 打印环境
		fmt.Println("enviroment:", envService.AppEnv())
	},
}

var envListCommand = &cobra.Command{
	Use:        "list",
	Cron:       nil,
	CronSpecs:  nil,
	Aliases:    nil,
	SuggestFor: nil,
	Short:      "获取所有环境变量",

	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		// 获取env 环境
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		envs := envService.All()
		outs := [][]string{}

		for k, v := range envs {
			outs = append(outs, []string{k, v})
		}
		util.PrettyPrint(outs)
	},
}
