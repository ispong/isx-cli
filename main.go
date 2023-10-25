package main

import (
	"isxcode.com/isxcode/isx-cli/src/action"
	"isxcode.com/isxcode/isx-cli/src/utils"
	"os"
)

func main() {

	utils.CheckCommandArgs()

	actionCode := os.Args[1]
	switch actionCode {
	case "reset":
		action.Reset()
	case "login":
		action.Login()
	case "logout":
		action.Reset()
	case "clone":
		action.Reset()
	//case "list":
	//	action.ListPackage()
	case "choose":
		action.Reset()
	case "show":
		action.Reset()
	case "idea":
		action.Reset()
	case "vscode":
		action.Reset()
	case "clean":
		action.Reset()
	case "start":
		action.Reset()
	case "package":
		action.Reset()
	case "docker":
		action.Reset()
	case "deploy":
		action.Reset()
	case "website":
		action.Reset()
	case "git":
		action.Reset()
	case "get":
		action.Reset()
	case "pr":
		action.Reset()
	case "branch":
		action.Reset()
	case "frontend":
		action.Reset()
	case "backend":
		action.Reset()
	case "web":
		action.Reset()
	case "home":
		action.Reset()
	case "install":
		action.Reset()
	case "remove":
		action.Reset()
	case "format":
		action.Reset()
	case "version":
		action.Reset()
	}
}
