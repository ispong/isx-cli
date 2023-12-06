package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "提交博客",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("提交代码")
		executeCommand := "cd /Users/ispong/code/ispong-blogs && git add . && git commit -m \":memo: 写博客\" && git push origin main"
		result := exec.Command("bash", "-c", executeCommand)
		result.Stdout = os.Stdout
		result.Stderr = os.Stderr

		err := result.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}
