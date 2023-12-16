/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
)

var (
	homeProjectName string
)

// homeCmd represents the home command
var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "快速进入项目目录，举例：isx home",
	Long:  `快速进入项目目录，举例：isx home`,
	Run: func(cmd *cobra.Command, args []string) {

		homeProjectName = viper.GetString("current-project.name")
		executeCommand := "cd " + viper.GetString(homeProjectName+".dir") + "/" + viper.GetString(homeProjectName+".name")
		cloneCmd := exec.Command("bash", "-c", executeCommand)
		cloneCmd.Stdout = os.Stdout
		cloneCmd.Stderr = os.Stderr
		err := cloneCmd.Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		} else {
			fmt.Println(viper.GetString(homeProjectName+".dir") + "/" + viper.GetString(homeProjectName+".name") + ":  跳转成功")
		}
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
