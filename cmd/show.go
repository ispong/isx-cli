package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "展示当前项目信息",
	Long:  `isx show`,
	Run: func(cmd *cobra.Command, args []string) {
		showCmdMain()
	},
}

func showCmdMain() {
	fmt.Println("当前开发项目")
	projectName := viper.GetString("current-project.name")
	fmt.Println("名称：" + viper.GetString(projectName+".name"))
	fmt.Println("描述：" + viper.GetString(projectName+".describe"))
	fmt.Println("路径：" + viper.GetString(projectName+".dir") + "/" + viper.GetString(projectName+".name"))
}
