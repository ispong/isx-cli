package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(packageCmd)
}

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: printCommand("isx package", 65) + "| 使用gradle打包",
	Long:  `isx package`,
	Run: func(cmd *cobra.Command, args []string) {
		packageCmdMain()
	},
}

func packageCmdMain() {

	projectName := viper.GetString("current-project.name")
	projectPath := viper.GetString(projectName+".dir") + "/" + projectName

	gradleCmd := exec.Command("./gradlew", "clean", "package")
	gradleCmd.Stdout = os.Stdout
	gradleCmd.Stderr = os.Stderr
	gradleCmd.Dir = projectPath
	err := gradleCmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("执行成功")
	}
}
