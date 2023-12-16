/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// chooseCmd represents the choose command
var chooseCmd = &cobra.Command{
	Use:   "choose",
	Short: "选择当前开发项目，举例：isx choose",
	Long:  `从isxcode组织中选择开发项目`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choose called")
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
