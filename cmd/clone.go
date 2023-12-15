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

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "下载项目代码，举例：isx clone",
	Long:  `下载项目代码，举例：isx clone`,
	Run: func(cmd *cobra.Command, args []string) {
		// 输入项目编号
		inputProjectNumber()
		// 输入下载路径
		inputProjectPath()
		// 下载主项目代码
		cloneCode()
		// 保存配置
		saveConfig()
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Project struct {
	Name       string   `yaml:"name"`
	Number     int      `yaml:"number"`
	Describe   string   `yaml:"describe"`
	Dir        string   `yaml:"dir"`
	Repository string   `yaml:"repository"`
	Sub        []string `yaml:"sub"`
}

type Config struct {
	Account string `yaml:"account"`
	//Projects map[string]Project `yaml:"projects"`
	Projects []Project `yaml:"projects"`
	Token    string    `yaml:"token"`
	Current  int       `yaml:"current"`
}

var (
	config        Config
	projectNumber int
	projectPath   string
)

func inputProjectNumber() {
	fmt.Println("请输入下载项目编号：")
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 打印可选项目编号
	for index, project := range config.Projects {
		fmt.Println("[" + strconv.Itoa(index) + "] " + project.Name + ": " + project.Describe)
	}
	// 输入编号
	_, err = fmt.Scanln(&projectNumber)
}

func inputProjectPath() {

	fmt.Print("请输入安装路径:")
	_, err := fmt.Scanln(&projectPath)
	if err != nil {
		return
	}

	_, err = os.Stat(projectPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(projectPath, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func cloneCode() {

	mainRepository := config.Projects[projectNumber].Repository
	strings.Replace(mainRepository, "isxcode", config.Account, -1)
	strings.Replace(mainRepository, "https://", "https://"+config.Token+"@", -1)

	// 下载主项目代码
	gitCmd := exec.Command("git", "clone", mainRepository)
	gitCmd.Stdout = os.Stdout
	gitCmd.Stderr = os.Stderr
	gitCmd.Dir = projectPath
	err := gitCmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("下载成功")
	}

	// 下载子项目代码
	subRepository := config.Projects[projectNumber].Sub
	for _, repository := range subRepository {
		strings.Replace(repository, "isxcode", config.Account, -1)
		strings.Replace(repository, "https://", "https://"+config.Token+"@", -1)
		gitCmd := exec.Command("git", "clone", repository)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		gitCmd.Dir = projectPath + "/" + config.Projects[projectNumber].Name
		err := gitCmd.Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		} else {
			fmt.Println("下载成功")
		}
	}
}

func saveConfig() {
	config.Projects[projectNumber].Dir = projectPath
	viper.Set("projects", config.Projects)
	viper.Set("current", projectNumber)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
