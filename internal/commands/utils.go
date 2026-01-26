package commands

import "os"

func IsRepoInitialized() bool {
	_, err := os.Stat(".loki")
	return err == nil
}
