package utils

import (
	"fmt"
	"os"
)

func CheckCommandArgs() {

	if len(os.Args) < 2 {

		commandFormat := "%-40s"
		fmt.Println("请输入完整命令，参考如下：")
		fmt.Println(fmt.Sprintf(commandFormat, "isx reset") + "-- 重置配置")
		fmt.Println(fmt.Sprintf(commandFormat, "isx login") + "-- 用户登录")
		fmt.Println(fmt.Sprintf(commandFormat, "isx logout") + "-- 用户退出")
		fmt.Println(fmt.Sprintf(commandFormat, "isx clone") + "-- 下载项目")
		fmt.Println(fmt.Sprintf(commandFormat, "isx list") + "-- 查看本地项目列表")
		fmt.Println(fmt.Sprintf(commandFormat, "isx choose") + "-- 选择开发项目")
		fmt.Println(fmt.Sprintf(commandFormat, "isx branch <issue_number>") + "-- 切开发分支")
		fmt.Println(fmt.Sprintf(commandFormat, "isx install") + "-- 安装项目依赖")
		fmt.Println(fmt.Sprintf(commandFormat, "isx start") + "-- 启动项目")
		fmt.Println(fmt.Sprintf(commandFormat, "isx package") + "-- 打包项目")
		fmt.Println(fmt.Sprintf(commandFormat, "isx format") + "-- 格式化代码")
		fmt.Println(fmt.Sprintf(commandFormat, "isx get <issue_number>") + "-- 获取需求分支名称")
		fmt.Println(fmt.Sprintf(commandFormat, "isx <git_command>") + "-- 在所有模块中执行git命令")
		fmt.Println(fmt.Sprintf(commandFormat, "isx pr <issue_number> '<pr_command>'") + "-- 提交所有模块的pr")
		os.Exit(0)
	}
}
