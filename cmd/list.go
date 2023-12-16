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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "陈列项目列表，举例：isx list",
	Long:  `陈列项目列表，举例：isx list`,
	Run: func(cmd *cobra.Command, args []string) {
		// 打印项目列表
		projectList := viper.GetStringSlice("project-list")
		for index, projectName := range projectList {
			fmt.Println("[" + strconv.Itoa(index) + "] " + viper.GetString(projectName+".name") + ": " + viper.GetString(projectName+".describe") + " 下载状态 【" + viper.GetString(projectName+".repository.download") + "】")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
