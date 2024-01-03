package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "陈列项目列表",
	Long:  `isx list`,
	Run: func(cmd *cobra.Command, args []string) {
		listCmdMain()
	},
}

func listCmdMain() {

	projectList := viper.GetStringSlice("project-list")
	for index, projectName := range projectList {
		fmt.Println("[" + strconv.Itoa(index) + "] " + viper.GetString(projectName+".name") + ": " + viper.GetString(projectName+".describe") + " 下载状态 【" + viper.GetString(projectName+".repository.download") + "】")
	}
}
