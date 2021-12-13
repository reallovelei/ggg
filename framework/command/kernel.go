package command

import (
	"github.com/reallovelei/ggg/framework/cobra"
)

func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(initAppCommand())
	root.AddCommand(initCronCommand())
}
