/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	showProjectName string
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "展示当前项目信息，举例：isx show",
	Long:  `展示当前项目信息，举例：isx show`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("当前开发项目信息")
		showProjectName = viper.GetString("current-project.name")
		fmt.Println("名称：" + viper.GetString(showProjectName+".name"))
		fmt.Println("描述：" + viper.GetString(showProjectName+".describe"))
		fmt.Println("路径：" + viper.GetString(showProjectName+".dir") + "/" + viper.GetString(showProjectName+".name"))
		fmt.Println("下载状态：" + viper.GetString(showProjectName+".repository.download"))
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
