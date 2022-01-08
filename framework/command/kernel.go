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
	// root.AddCommand(initConfigCommand())
}
