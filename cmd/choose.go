/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

// chooseCmd represents the choose command
var chooseCmd = &cobra.Command{
	Use:   "choose",
	Short: "选择当前开发项目，举例：isx choose",
	Long:  `从isxcode组织中选择开发项目`,
	Run: func(cmd *cobra.Command, args []string) {

		// 打印项目列表
		projectList := viper.GetStringSlice("project-list")
		for index, projectName := range projectList {
			fmt.Println("[" + strconv.Itoa(index) + "] " + viper.GetString(projectName+".name") + ": " + viper.GetString(projectName+".describe") + " 下载状态 【" + viper.GetString(projectName+".repository.download") + "】")
		}
		fmt.Println("请输入下载项目编号：")
		fmt.Scanln(&projectNumber)
		projectName = projectList[projectNumber]
		fmt.Println("切换到项目：" + projectName)

		// 将当前的项目设置
		viper.Set("current-project.name", projectList[projectNumber])
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(chooseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chooseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chooseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
