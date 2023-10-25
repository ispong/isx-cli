package main

import (
	action2 "isxcode.com/isxcode/isx-cli/action"
	"isxcode.com/isxcode/isx-cli/utils"
	"os"
)

func main() {

	utils.CheckCommandArgs()

	actionCode := os.Args[1]
	switch actionCode {
	case "reset":
		action2.Reset()
	case "login":
		action2.Login()
	case "logout":
		action2.Reset()
	case "clone":
		action2.Reset()
	//case "list":
	//	action.ListPackage()
	case "choose":
		action2.Reset()
	case "show":
		action2.Reset()
	case "idea":
		action2.Reset()
	case "vscode":
		action2.Reset()
	case "clean":
		action2.Reset()
	case "start":
		action2.Reset()
	case "package":
		action2.Reset()
	case "docker":
		action2.Reset()
	case "deploy":
		action2.Reset()
	case "website":
		action2.Reset()
	case "git":
		action2.Reset()
	case "get":
		action2.Reset()
	case "pr":
		action2.Reset()
	case "branch":
		action2.Reset()
	case "frontend":
		action2.Reset()
	case "backend":
		action2.Reset()
	case "web":
		action2.Reset()
	case "home":
		action2.Reset()
	case "install":
		action2.Reset()
	case "remove":
		action2.Reset()
	case "format":
		action2.Reset()
	case "version":
		action2.Reset()
	}
}
