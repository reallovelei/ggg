package command

import (
    "fmt"
    "github.com/reallovelei/ggg/framework/cobra"
    "github.com/reallovelei/ggg/framework/contract"
)

var ExampleCommand = &cobra.Command{
    Use:"example",
    Short:"example for framework",
    Run: func(c *cobra.Command, args []string) {
        container := c.GetContainer()
        appService := container.MustMake(contract.AppKey).(contract.App)
        fmt.Println("app base path:", appService.BasePath())
    },
}
