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

// gradleCmd represents the gradle command
var gradleCmd = &cobra.Command{
	Use:   "gradle",
	Short: "在项目内执行gradle命令，举例：isx gradle [install|start|clean|format|package]",
	Long:  `gradle install、gradle start、gradle clean、gradle format`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName = viper.GetString("current-project.name")
		projectPath = viper.GetString(projectName + ".dir")

		gradleCmd := exec.Command("./gradlew", args...)
		gradleCmd.Stdout = os.Stdout
		gradleCmd.Stderr = os.Stderr
		gradleCmd.Dir = projectPath + "/" + projectName
		err := gradleCmd.Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		} else {
			fmt.Println("执行成功")
		}
	},
}

func init() {
	rootCmd.AddCommand(gradleCmd)
}
