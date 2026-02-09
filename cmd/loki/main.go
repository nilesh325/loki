package main

import (
	"fmt"
	"os"
	"path/filepath"

	"loki/internal/commands"
	"loki/internal/core"
	"loki/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		commands.Help()
		return
	}

	cwd, _ := os.Getwd()
	absPath, _ := filepath.Abs(cwd)
	if os.Args[0] == "./loki" && os.Args[1] != "help" && os.Args[1] != "init" {
		_, check := core.IsRepoInitialized(absPath)
		if !check {
			fmt.Println(utils.ColorText("fatal:", "error") + utils.ColorText(string(cwd), "notice") + utils.ColorText(" not a loki repository \n(or any of the parent directories)", "error"))
			return
		}
	}

	switch os.Args[1] {
	case "init":
		commands.Init()
	case "add":
		commands.Add(os.Args[2:])
	case "commit":
		commands.Commit(os.Args[2:])
	case "status":
		commands.Status()
	case "log":
		commands.Log()
	case "help":
		commands.Help()
	case "config":
		commands.Config(os.Args[2:])
	default:
		fmt.Println("Unknown command:", os.Args[1])
		commands.Help()
	}
}
