/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gradleCmd represents the gradle command
var gradleCmd = &cobra.Command{
	Use:   "gradle",
	Short: "在项目内执行gradle命令，举例：isx gradle install",
	Long:  `gradle install、gradle start、gradle clean、gradle format`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gradle called")
	},
}

func init() {
	rootCmd.AddCommand(gradleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gradleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gradleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
