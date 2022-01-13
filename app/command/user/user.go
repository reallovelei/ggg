package user

import (
	"fmt"
	"github.com/reallovelei/ggg/framework/cobra"
)

var UserCommand = &cobra.Command{
	Use:   "user",
	Short: "user",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		fmt.Println(container)
		return nil
	},
}
