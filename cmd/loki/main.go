package main

import (
	"fmt"
	"os"

	"loki/internal/commands"
)

func main() {
	if len(os.Args) >= 2 && os.Args[1] != "init" && os.Args[1] != "help" {
		if !commands.IsRepoInitialized() {
			fmt.Println("fatal: not a loki repository (or any of the parent directories): .loki")
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
	default:
		fmt.Println("Unknown command:", os.Args[1])
		commands.Help()
	}
}
