package command

import (
	"fmt"
	"github.com/reallovelei/ggg/framework/cobra"
	"log"
	"os/exec"
)

func initBuildCommand() *cobra.Command {
	buildCommand.AddCommand(buildFrontendCommand)
	buildCommand.AddCommand(buildBackendCommand)
	buildCommand.AddCommand(buildAllCommand)
	return buildCommand
}

var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "编译相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

var buildBackendCommand = &cobra.Command{
	Use:   "backend",
	Short: "使用go编译后端",
	RunE: func(c *cobra.Command, args []string) error {
		path, err := exec.LookPath("go")
		if err != nil {
			log.Fatalln("ggg go: please install go in path first")
		}

		cmd := exec.Command(path, "build", "-o", "ggg", "./")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("go build error:")
			fmt.Println(string(out))
			fmt.Println("--------------")
			return err
		}
		fmt.Println("build success please run ./ggg direct")
		return nil
	},
}

var buildFrontendCommand = &cobra.Command{
	Use:   "frontend",
	Short: "使用npm编译前端",
	RunE: func(c *cobra.Command, args []string) error {
		// 获取path路径下的npm命令
		path, err := exec.LookPath("npm")
		if err != nil {
			log.Fatalln("请安装npm在你的PATH路径下")
		}

		// 进到前端项目目录
		cmd := exec.Command("cd", "./front")
		errCd := cmd.Run()
		if errCd != nil {
			fmt.Println("cd error", errCd)
		}

		// 执行npm run build
		cmd = exec.Command(path, "run", "build")
		// 将输出保存在out中
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("=============  前端编译失败 ============")
			fmt.Println(string(out))
			fmt.Println("=============  前端编译失败 ============")
			return err
		}
		// 打印输出
		fmt.Print(string(out))
		fmt.Println("=============  前端编译成功 ============")
		return nil
	},
}

var buildAllCommand = &cobra.Command{
	Use:   "all",
	Short: "同时编译前端和后端",
	RunE: func(c *cobra.Command, args []string) error {
		err := buildFrontendCommand.RunE(c, args)
		if err != nil {
			return err
		}
		err = buildBackendCommand.RunE(c, args)
		if err != nil {
			return err
		}
		return nil
	},
}
