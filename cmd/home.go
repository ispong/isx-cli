/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// homeCmd represents the home command
var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "快速进入项目目录，举例：isx home",
	Long:  `快速进入项目目录，举例：isx home`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("home called")
	},
}

func init() {
	rootCmd.AddCommand(homeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// homeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// homeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
