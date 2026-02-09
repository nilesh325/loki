package commands

import (
	"fmt"
	"loki/internal/config"
	"os"
	"strings"
)

func Config(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: loki config [--local|--global|--system] key [value]")
		return
	}

	level := "local" // default
	key := ""
	value := ""
	for _, arg := range args {
		if arg == "--local" || arg == "--global" || arg == "--system" {
			level = strings.TrimPrefix(arg, "--")
		} else if key == "" {
			key = arg
		} else {
			value = arg
		}
	}

	repoRoot := config.FindRepoRoot(os.Getenv("PWD"))
	cfg := config.NewConfig()
	cfg.Load(repoRoot)

	if value == "" {
		// Get
		v := cfg.Get(key)
		if v == "" {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, v)
		}
	} else {
		// Set
		err := cfg.Set(level, repoRoot, key, value)
		if err != nil {
			fmt.Printf("Error setting %s: %v\n", key, err)
		} else {
			fmt.Printf("Set %s=%s (%s)\n", key, value, level)
		}
	}
}
