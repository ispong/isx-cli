package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"os/user"
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "编译本地代码",
	Long:  `isx build`,
	Run: func(cmd *cobra.Command, args []string) {
		buildCmdMain()
	},
}

func buildCmdMain() {

	projectName := viper.GetString("current-project.name")
	projectPath := viper.GetString(projectName+".dir") + "/" + viper.GetString(projectName+".name")
	buildImage := "registry.cn-shanghai.aliyuncs.com/isxcode/zhiqingyun-build"
	usr, _ := user.Current()

	// 获取gradle缓存目录
	cacheGradleDir := viper.GetString("cache.gradle.dir")
	if cacheGradleDir == "" {
		cacheGradleDir = usr.HomeDir + "/.gradle"
	}
	_, err := os.Stat(cacheGradleDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(cacheGradleDir, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// 获取pnpm缓存目录
	cachePnpmDir := viper.GetString("cache.pnpm.dir")
	if cachePnpmDir == "" {
		cachePnpmDir = usr.HomeDir + "/.pnpm-store"
	}
	_, err = os.Stat(cachePnpmDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(cachePnpmDir, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// 下载主项目代码
	buildCommand := "docker run " +
		"--rm " +
		"-v " + projectPath + ":/spark-yun " +
		"-v " + cachePnpmDir + ":/root/.pnpm-store " +
		"-v " + cacheGradleDir + ":/root/.gradle" +
		"-d " + buildImage
	buildCmd := exec.Command("bash", "-c", buildCommand)
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	err = buildCmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("代码正在编译")
	}
}
