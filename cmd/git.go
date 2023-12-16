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

// gitCmd represents the git command
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "在项目内执行git命令，举例：isx git <git command>",
	Long:  `在项目内执行git命令，举例：isx git <git command>`,
	Run: func(cmd *cobra.Command, args []string) {

		projectName = viper.GetString("current-project.name")
		projectPath = viper.GetString(projectName + ".dir")

		// 进入主项目执行命令
		gitCmd := exec.Command("git", args...)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		gitCmd.Dir = projectPath + "/" + projectName
		err := gitCmd.Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		} else {
			fmt.Println("执行成功")
		}

		// 进入子项目执行命令
		var subRepository []Repository
		viper.UnmarshalKey(projectName+".sub-repository", &subRepository)
		for _, repository := range subRepository {

			gitCmd := exec.Command("git", args...)
			gitCmd.Stdout = os.Stdout
			gitCmd.Stderr = os.Stderr
			gitCmd.Dir = projectPath + "/" + projectName + "/" + repository.Name
			err := gitCmd.Run()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			} else {
				fmt.Println("执行成功")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)
}
