package command

import (
	"github.com/reallovelei/ggg/framework/cobra"
)

func AddKernelCommands(root *cobra.Command) {
	// app Command
	root.AddCommand(initAppCommand())
	// env
	root.AddCommand(initEnvCommand())
	// cron Command
	root.AddCommand(initCronCommand())
	// config
	root.AddCommand(initConfigCommand())
	// build命令
	root.AddCommand(initBuildCommand())

	// cmd
	root.AddCommand(initCmdCommand())

	// middleware
	root.AddCommand(initMiddlewareCommand())

	// new
	root.AddCommand(initNewCommand())

	// provider一级命令
	root.AddCommand(initProviderCommand())

}
