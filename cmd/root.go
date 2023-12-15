/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "isx",
	Short: "cli for isxcode app",
	Long: `
 ____ _____ __ __           __  _      ____ 
|    / ___/|  |  |         /  ]| |    |    |
 |  (   \_ |  |  | _____  /  / | |     |  | 
 |  |\__  ||_   _||     |/  /  | |___  |  | 
 |  |/  \ ||     ||_____/   \_ |     | |  | 
 |  |\    ||  |  |      \     ||     | |  | 
|____|\___||__|__|       \____||_____||____|

欢迎使用isx-cli脚手架

`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)

	// 解析配置文件
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.isx/isx-config.yml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {

	// 获取home目录
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 初始化配置文件信息
	viper.SetConfigFile(home + "/.isx/isx-config.yml")

	// 判断配置文件是否存在
	if err := viper.ReadInConfig(); err != nil {

		// 判断文件夹是否存在，不存在则新建
		_, err := os.Stat(home + "/.isx")
		if os.IsNotExist(err) {
			err := os.Mkdir(home+"/.isx", 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// 判断文件是否存在，不存在则新建
		_, err = os.Stat(home + "/.isx/isx-config.yml")
		if os.IsNotExist(err) {
			// 初始化配置
			viper.SetConfigType("yaml")
			var yamlExample = []byte(`
account: ''
token: ''
current: ''
projects:
  -   name: flink-yun
      number: 0
      dir: ''
      describe: 至爻云-流（至流云）
      repository: https://github.com/isxcode/flink-yun.git
      sub:
        - https://github.com/isxcode/flink-yun-vip.git
  -   name: spark-yun
      number: 1
      dir: ''
      describe: 至爻云-轻（至轻云）
      repository: https://github.com/isxcode/spark-yun.git
      sub:
        - https://github.com/isxcode/spark-yun-vip.git
`)
			err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// 持久化配置
			err = viper.SafeWriteConfigAs(home + "/.isx/isx-config.yml")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
