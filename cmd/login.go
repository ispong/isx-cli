/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "github用户登录",
	Long:  `github用户登录`,
	Run: func(cmd *cobra.Command, args []string) {
		var account string
		var token string
		fmt.Print("请输入github账号:")
		_, err := fmt.Scanln(&account)
		if err != nil {
			return
		}
		fmt.Println("快捷链接：https://github.com/settings/tokens")
		fmt.Print("请输入token:")
		_, err = fmt.Scanln(&token)
		if err != nil {
			return
		}
		// 检查token是否有效
		checkGithubToken(account, token)
		// 保存配置
		saveAccountAndToken(account, token)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.
	//loginCmd.Flags().StringP("account", "a", "", "github账号")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")
}

// 保存配置
func saveAccountAndToken(account string, token string) {
	viper.Set("account", account)
	viper.Set("token", token)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 检测token是否合法
func checkGithubToken(account string, token string) {

	headers := http.Header{}
	headers.Set("Accept", "application/vnd.github+json")
	headers.Set("Authorization", "Bearer "+token)
	headers.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/octocat", nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		os.Exit(1)
	}

	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		os.Exit(1)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 解析结果
	if resp.StatusCode == http.StatusOK {
		fmt.Println("登录成功，欢迎使用isx-cli开发工具")
	} else {
		fmt.Println("无法验证token合法性，登录失败")
		os.Exit(0)
	}
}
