/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import "C"
import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	projectNumber int
	projectPath   string
	projectName   string
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "下载项目代码，举例：isx clone",
	Long:  `下载项目代码，举例：isx clone`,
	Run: func(cmd *cobra.Command, args []string) {

		// 输入项目编号
		inputProjectNumber()

		// 输入安装路径
		inputProjectPath()

		// 下载项目代码
		cloneProjectCode()

		// 将默认项目设置为当前选项
		viper.Set("current-project.name", projectName)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}

func inputProjectNumber() {

	// 打印可选项目编号
	projectList := viper.GetStringSlice("project-list")
	for index, projectName := range projectList {
		fmt.Println("[" + strconv.Itoa(index) + "] " + viper.GetString(projectName+".name") + ": " + viper.GetString(projectName+".describe"))
	}
	fmt.Println("请输入下载项目编号：")

	// 输入项目编号
	fmt.Scanln(&projectNumber)
	projectName = projectList[projectNumber]
}

func inputProjectPath() {

	fmt.Print("请输入安装路径:")
	fmt.Scanln(&projectPath)

	// 目录不存在则报错
	_, err := os.Stat(projectPath)
	if os.IsNotExist(err) {
		fmt.Println("目录不存在，请重新输入")
		os.Exit(1)
	}

	// 保存配置
	viper.Set(projectName+".dir", projectPath)
	viper.WriteConfig()
}

type Repository struct {
	Download string `yaml:"download"`
	Url      string `yaml:"url"`
}

func cloneProjectCode() {

	// 构建下载链接
	projectName := viper.GetStringSlice("project-list")[projectNumber]
	mainRepository := viper.GetString(projectName + ".repository.url")
	strings.Replace(mainRepository, "isxcode", viper.GetString("user.account"), -1)
	strings.Replace(mainRepository, "https://", "https://"+viper.GetString("user.token")+"@", -1)

	// 下载主项目代码
	executeCommand := "git clone " + mainRepository
	cloneCmd := exec.Command("bash", "-c", executeCommand)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	cloneCmd.Dir = projectPath
	err := cloneCmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("下载成功")
		viper.Set(projectName+".repository.download", "ok")
		viper.WriteConfig()
	}

	// 下载子项目代码
	var subRepository []Repository
	viper.UnmarshalKey(projectName+".sub-repository", &subRepository)
	for index, repository := range subRepository {
		metaRepository := repository.Url
		strings.Replace(metaRepository, "isxcode", viper.GetString("user.account"), -1)
		strings.Replace(metaRepository, "https://", "https://"+viper.GetString("user.toke")+"@", -1)
		executeCommand := "git clone " + metaRepository
		cloneCmd := exec.Command("bash", "-c", executeCommand)
		cloneCmd.Stdout = os.Stdout
		cloneCmd.Stderr = os.Stderr
		cloneCmd.Dir = projectPath + "/" + projectName
		err := cloneCmd.Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		} else {
			fmt.Println("下载成功")
			subRepository[index].Download = "ok"
			viper.Set(projectName+".sub-repository", subRepository)
			viper.WriteConfig()
		}
	}
}
