package command

import (
	"github.com/reallovelei/ggg/app/command/demo"
    "github.com/reallovelei/ggg/app/command/user"
    "github.com/reallovelei/ggg/framework"
	"github.com/reallovelei/ggg/framework/cobra"
	"github.com/reallovelei/ggg/framework/command"
)

func RunCommand(container framework.Container) error {

	var rootCmd = &cobra.Command{
		Use:   "ggg",
		Short: "ggg 命令",
		Long:  "ggg 框架的命令行工具，，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现cobra默认的completion子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	// 为根Command设置服务容器
	rootCmd.SetContainer(container)

	// 绑定框架的命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)

	// 执行RootCommand
	return rootCmd.Execute()
}

// 绑定业务的命令
func AddAppCommand(rootCmd *cobra.Command) {
	//  demo 例子
	rootCmd.AddCommand(demo.InitFoo())
    rootCmd.AddCommand(user.UserCommand)
    rootCmd.AddCommand(demo.Foo1Command)
	rootCmd.AddCronCommand("* * * * * *", demo.FooCommand)
	rootCmd.AddCronCommand("* * * * * *", demo.Foo1Command)
}
