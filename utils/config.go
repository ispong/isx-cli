package utils

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
	"path/filepath"
)

var IsxConfigTemp = "" +
	"develop-project: spark-yun\n" +
	"projects:\n" +
	"  flink-yun:\n" +
	"    describe: 至流云 - 基于flink的大数据平台\n" +
	"    dir: ''\n" +
	"    repository: https://github.com/isxcode/flink-yun.git\n" +
	"    sub-repository:\n" +
	"    - https://github.com/isxcode/flink-yun-vip.git\n" +
	"  isx-base:\n" +
	"    describe: isxcode模版\n" +
	"    dir: ''\n" +
	"    repository: https://github.com/isxcode/isx-base.git\n" +
	"    sub-repository:\n" +
	"    - https://github.com/isxcode/isx-base-vip.git\n" +
	"  isx-cli:\n" +
	"    describe: isx-cli 脚手架\n" +
	"    dir: ''\n" +
	"    repository: https://github.com/isxcode/isx-cli.git\n" +
	"    sub-repository: []\n" +
	"  spark-yun:\n" +
	"    describe: \"至轻云 - 基于spark的大数据平台\n" +
	"    dir: ''\n" +
	"    repository: https://github.com/isxcode/spark-yun.git\n" +
	"    sub-repository:\n" +
	"    - https://github.com/isxcode/spark-yun-vip.git\n" +
	"user:\n" +
	"  account: ''\n" +
	"  token: ''"

func getIsxConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".isx", "isx-config.yml")
}

func ClearConfig() {
	err := os.Remove(getIsxConfigPath())
	if err != nil {
		fmt.Println("清楚配置失败")
		os.Exit(1)
	}
}

func initIsxConfig() {

	homeDir, _ := os.UserHomeDir()

	isxConfigDir := filepath.Join(homeDir, ".isx")
	if _, err := os.Stat(isxConfigDir); err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(isxConfigDir, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	isxConfigPath := filepath.Join(homeDir, ".isx", "isx-config.yml")
	if _, err := os.Stat(isxConfigPath); err != nil && os.IsNotExist(err) {
		f, err := os.OpenFile(isxConfigPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}(f)

		yamlEncoder := yaml.NewEncoder(f)
		err = yamlEncoder.Encode(IsxConfigTemp)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	//if err != nil {
	//	return Config{}, err
	//}
	//defer configFile.Close()
	//
	//// 读取配置文件的内容
	//configBytes, err := ioutil.ReadAll(configFile)
	//if err != nil {
	//	return Config{}, err
	//}
	//
	//// 解析 YAML 配置文件
	//var config Config
	//err = yaml.Unmarshal(configBytes, &config)
	//if err != nil {
	//	return Config{}, err
	//}
}

func GetIsxConfig() {
	initIsxConfig()
}
