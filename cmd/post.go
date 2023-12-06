package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(postCmd)
}

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "新建博客",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("新建博客")

		titleFirst := args[0]
		titleLast := args[1]
		folder := ""

		// 知识积累方面
		githubList := []string{"docsify", "git", "github", "hexo", "markdown", "vscode"}
		for i := range githubList {
			if titleFirst == githubList[i] {
				folder = "github"
				break
			}
		}

		// 大数据相关
		hadoopList := []string{"hadoop", "hbase", "hive", "flink", "spark", "clickhouse", "doris", "kafka", "sqoop", "canal", "zookeeper", "atlas", "cdh", "solr"}
		for i := range hadoopList {
			if titleFirst == hadoopList[i] {
				folder = "hadoop"
				break
			}
		}

		// 云原生相关
		kubernetesList := []string{"go", "golang", "kubernetes", "docker", "rancher", "jenkins"}
		for i := range kubernetesList {
			if titleFirst == kubernetesList[i] {
				folder = "kubernetes"
				break
			}
		}

		// 操作系统相关
		osList := []string{"linux", "mac", "windows", "ngrok", "clash"}
		for i := range osList {
			if titleFirst == osList[i] {
				folder = "os"
				break
			}
		}

		// ai智能相关
		pytorchList := []string{"anaconda", "pytorch", "python", "pycharm", "scrapy"}
		for i := range pytorchList {
			if titleFirst == pytorchList[i] {
				folder = "pytorch"
				break
			}
		}

		// 后端开发相关
		springList := []string{"java", "spring", "idea", "gradle", "maven", "rabbitmq", "redis"}
		for i := range springList {
			if titleFirst == springList[i] {
				folder = "spring"
				break
			}
		}

		// 前端开发相关
		vueList := []string{"node", "typescript", "vue", "webstorm", "vite", "nginx", "html", "sass", "antdesign", "element"}
		for i := range vueList {
			if titleFirst == vueList[i] {
				folder = "vue"
				break
			}
		}

		// 数据源相关
		dbList := []string{"mongodb", "mysql", "oracle", "sqlserver", "sqlserver", "postgre", "h2"}
		for i := range dbList {
			if titleFirst == dbList[i] {
				folder = "db"
				break
			}
		}

		//
		if folder == "" {
			fmt.Println("该分类不支持")
			return
		}

		executeCommand := "cd /Users/ispong/code/ispong-blogs && hexo new " + titleFirst + " -p " + folder + "/" + titleFirst + "/\"" + titleFirst + " " + titleLast + "\"" + " " + "\"" + titleFirst + " " + titleLast + "\""
		fmt.Println(executeCommand)
		result := exec.Command("bash", "-c", executeCommand)
		result.Stdout = os.Stdout
		result.Stderr = os.Stderr

		err := result.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("创建完成")
	},
}
