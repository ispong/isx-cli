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

type Member struct {
	Account string `json:"login"`
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

func checkoutCmdMain(branchName string) {

	// 本地有分支 直接切换
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
		checkoutRemoteBranch(projectPath, branch)

		var subRepository []Repository
		viper.UnmarshalKey(viper.GetString("current-project.name")+".sub-repository", &subRepository)
		for _, repository := range subRepository {
			checkoutRemoteBranch(projectPath+"/"+repository.Name, branch)
		}

		return
	}

	// 本地没有分支，远程没有分支，找成员的远程是否有分支，直接切换
	memberHasBranch := printMember(branchName)
	if memberHasBranch {

		// 输入成员名字
		fmt.Print("请输入从哪个成员出拉取分支:")
		var memberName string
		fmt.Scanln(&memberName)

	}

	//// 判断远程仓库是否已添加
	//repository := checkoutMemberRepository(memberName)
	//if repository == "" {
	//	// 添加成员的仓库
	//	// git remote add
	//	// git fetch memberName
	//	// git checkout branchNum memberName/branchNum
	//} else {
	//	// git fetch memberName
	//	// git checkout branchNum memberName/branchNum
	//}
	//
	//// 从isxcode远程切一个分支
	//
	//// 如果都没有，则去isxcode的latest分支中获取
	//
	//// 推到用户自己的远程
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

func printMember(branchNum string) bool {

	headers := http.Header{}
	headers.Set("Accept", "application/vnd.github+json")
	headers.Set("Authorization", "Bearer "+viper.GetString("user.token"))
	headers.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/orgs/isxcode/teams/spark-yun/members", nil)

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

	flag := false

	// 解析结果
	if resp.StatusCode == http.StatusOK {
		var people []Member
		err := json.Unmarshal(body, &people)
		if err != nil {
			fmt.Println("解析 JSON 失败:", err)
		}
		// 打印列表对象
		for _, person := range people {
			metaBranch := getGithubBranch(branchNum, person.Account)
			if metaBranch != "" {
				metaBranch = metaBranch + " 已创建"
				fmt.Println(fmt.Sprintf("%-*s", 20, person.Account), metaBranch)
			}
			flag = true
		}
	} else {
		fmt.Println("无法验证token合法性，登录失败")
		os.Exit(0)
	}
	return flag
}

func checkoutMemberRepository(memberName string) string {

	projectName := viper.GetString("current-project.name")
	projectPath := viper.GetString(projectName+".dir") + "/" + projectName

	cmd := exec.Command("bash", "-c", "git remote -v")
	cmd.Dir = projectPath

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("执行命令失败:", err)
		return ""
	}

	repositorys := strings.Split(string(output), "\n")
	for _, repository := range repositorys {
		if strings.Contains(repository, memberName) {
			return memberName
		}
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

func checkoutRemoteBranch(path string, branchName string) {

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
