/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")

		//mkdir -p ~/.pnpm-store
		//mkdir -p ~/.gradle
		//
		//isx build
		//isx run
		//
		//docker run \
		//    -v /Users/xiaoxiao/Downloads/spark-yun:/spark-yun \
		//    -v /Users/xiaoxiao/.pnpm-store:/root/.pnpm-store \
		//    -v /Users/xiaoxiao/.gradle:/root/.gradle \
		//    -d registry.cn-shanghai.aliyuncs.com/isxcode/zhiqingyun-build
		//
		//docker run \
		//    -v /Users/xiaoxiao/Downloads/spark-yun:/spark-yun \
		//    -e ENV_TYPE='ALL' \
		//    -p 8899:8080 \
		//    -d registry.cn-shanghai.aliyuncs.com/isxcode/zhiqingyun-local:latest

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
