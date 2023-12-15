/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "快速提交pr，举例：isx pr #123 '添加登录模块'",
	Long:  `快速提交pr，举例：isx pr #123 '添加登录模块'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pr called")
	},
}

func init() {
	rootCmd.AddCommand(prCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
