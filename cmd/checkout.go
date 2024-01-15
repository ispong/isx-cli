package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type GithubIssue struct {
	Body string `json:"body"`
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "创建需求分支",
	Long:  `isx checkout 123`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("使用方式不对，请重新输入命令")
			os.Exit(1)
		}

		checkoutCmdMain(args[0])
	},
}

func checkoutCmdMain(issueNumber string) {

	branchName := "GH-" + issueNumber

	// 本地有分支，直接切换
	branch := getLocalBranchName(branchName)
	if branch != "" {

		projectName := viper.GetString("current-project.name")
		projectPath := viper.GetString(projectName+".dir") + "/" + projectName
		checkoutLocalBranch(projectPath, branch)

		var subRepository []Repository
		viper.UnmarshalKey(viper.GetString("current-project.name")+".sub-repository", &subRepository)
		for _, repository := range subRepository {
			checkoutLocalBranch(projectPath+"/"+repository.Name, branch)
		}

		return
	}

	// 本地没有分支，远程有分支，直接切换
	branch = getGithubBranch(branchName, viper.GetString("user.account"))
	if branch != "" {

		projectName := viper.GetString("current-project.name")
		projectPath := viper.GetString(projectName+".dir") + "/" + projectName
		checkoutOriginBranch(projectPath, branch)

		var subRepository []Repository
		viper.UnmarshalKey(viper.GetString("current-project.name")+".sub-repository", &subRepository)
		for _, repository := range subRepository {
			checkoutOriginBranch(projectPath+"/"+repository.Name, branch)
		}

		return
	}

	// 远程没分支，isxcode仓库有分支，直接切换
	branch = getGithubBranch(branchName, "isxcode")
	if branch != "" {

		projectName := viper.GetString("current-project.name")
		projectPath := viper.GetString(projectName+".dir") + "/" + projectName
		checkoutUpstreamBranch(projectPath, branch)

		var subRepository []Repository
		viper.UnmarshalKey(viper.GetString("current-project.name")+".sub-repository", &subRepository)
		for _, repository := range subRepository {
			checkoutUpstreamBranch(projectPath+"/"+repository.Name, branch)
		}

		return
	}

	// 哪里都没有分支，自己创建分支
	fatherBranchName := getGithubIssueBranch(issueNumber)

	// 本地切出分支
	if fatherBranchName == "main" {
		projectName := viper.GetString("current-project.name")
		projectPath := viper.GetString(projectName+".dir") + "/" + projectName
		createMainBranch(projectPath, branch)

		var subRepository []Repository
		viper.UnmarshalKey(viper.GetString("current-project.name")+".sub-repository", &subRepository)
		for _, repository := range subRepository {
			createMainBranch(projectPath+"/"+repository.Name, branch)
		}

		return
	} else {
		projectName := viper.GetString("current-project.name")
		projectPath := viper.GetString(projectName+".dir") + "/" + projectName
		createReleaseBranch(projectPath, branch)

		var subRepository []Repository
		viper.UnmarshalKey(viper.GetString("current-project.name")+".sub-repository", &subRepository)
		for _, repository := range subRepository {
			createReleaseBranch(projectPath+"/"+repository.Name, branch)
		}

		return
	}
}

func getLocalBranchName(branchName string) string {

	projectName := viper.GetString("current-project.name")
	projectPath := viper.GetString(projectName+".dir") + "/" + projectName

	cmd := exec.Command("bash", "-c", "git branch -l "+"\""+branchName+"\"")
	cmd.Dir = projectPath

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("执行命令失败:", err)
		return ""
	}

	branches := strings.Split(string(output), "\n")
	for _, branch := range branches {
		branch = strings.ReplaceAll(strings.Replace(branch, "*", "", -1), " ", "")
		if branch == branchName {
			return branch
		}
	}

	return ""
}

func getGithubBranch(branchNum string, account string) string {

	headers := http.Header{}
	headers.Set("Accept", "application/vnd.github+json")
	headers.Set("Authorization", "Bearer "+viper.GetString("user.token"))
	headers.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+account+"/"+viper.GetString("current-project.name")+"/branches/"+branchNum, nil)

	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		os.Exit(1)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("关闭响应体失败:", err)
		}
	}(resp.Body)

	// 读取响应体内容
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		os.Exit(1)
	}

	// 解析结果
	if resp.StatusCode == http.StatusOK {
		return branchNum
	} else if resp.StatusCode == http.StatusNotFound {
		return ""
	} else {
		fmt.Println("无法验证token合法性，登录失败")
		os.Exit(1)
	}

	return ""
}

func checkoutLocalBranch(path string, branchName string) {

	// 下载主项目代码
	executeCommand := "git checkout " + branchName
	cloneCmd := exec.Command("bash", "-c", executeCommand)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	cloneCmd.Dir = path
	err := cloneCmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("本地存在" + branchName + "，切换成功")
	}
}

func createMainBranch(path string, branchName string) {

	executeCommand := "git fetch upstream && git checkout --track upstream/" + branchName
	cloneCmd := exec.Command("bash", "-c", executeCommand)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	cloneCmd.Dir = path
	err := cloneCmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("本地存在" + branchName + "，切换成功")
	}
}

func createReleaseBranch(path string, branchName string) {

	executeCommand := "git fetch upstream && git checkout --track upstream/" + branchName
	cloneCmd := exec.Command("bash", "-c", executeCommand)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	cloneCmd.Dir = path
	err := cloneCmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("本地存在" + branchName + "，切换成功")
	}
}

func checkoutOriginBranch(path string, branchName string) {

	executeCommand := "git fetch && git checkout --track origin/" + branchName
	cloneCmd := exec.Command("bash", "-c", executeCommand)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	cloneCmd.Dir = path
	err := cloneCmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("本地存在" + branchName + "，切换成功")
	}
}

func checkoutUpstreamBranch(path string, branchName string) {

	executeCommand := "git fetch && git checkout --track origin/" + branchName
	cloneCmd := exec.Command("bash", "-c", executeCommand)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	cloneCmd.Dir = path
	err := cloneCmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		fmt.Println("本地存在" + branchName + "，切换成功")
	}
}

func getGithubIssueBranch(issueNumber string) string {

	headers := http.Header{}
	headers.Set("Accept", "application/vnd.github+json")
	headers.Set("Authorization", "Bearer "+viper.GetString("user.token"))
	headers.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/isxcode/"+viper.GetString("current-project.name")+"/issues/"+issueNumber, nil)

	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		os.Exit(1)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("关闭响应体失败:", err)
		}
	}(resp.Body)

	// 读取响应体内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		os.Exit(1)
	}

	// 解析结果
	if resp.StatusCode == http.StatusOK {
		fmt.Println(string(body))
		var content GithubIssue
		err := json.Unmarshal(body, &content)
		if err != nil {
			fmt.Println("解析 JSON 失败:", err)
		}
		// 使用正则表达式查找匹配项
		versionStart := "### 版本号\n\nv"
		versionEnd := "\n\n### 缺陷内容"

		startIndex := strings.Index(content.Body, versionStart)
		endIndex := strings.Index(content.Body, versionEnd)

		if startIndex == -1 || endIndex == -1 {
			return "main"
		}

		version := content.Body[startIndex+len(versionStart) : endIndex]
		return version
	} else if resp.StatusCode == http.StatusNotFound {
		fmt.Println("issue不存在")
		os.Exit(1)
	} else {
		fmt.Println("无法验证token合法性，登录失败")
		os.Exit(1)
	}

	return ""
}
