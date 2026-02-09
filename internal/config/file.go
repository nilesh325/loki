package config

import (
	"os"
	"path/filepath"
)

func FindRepoRoot(start string) string {
	dir := start
	for {
		configPath := filepath.Join(dir, ".loki/config")
		if _, err := os.Stat(configPath); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}
