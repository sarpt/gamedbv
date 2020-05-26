package cmds

import (
	"os"
	"os/exec"
	"path"
)

const (
	defaultLocalPath = "."
)

func createPlatformArguments(platforms []string) []string {
	args := []string{}

	for _, platform := range platforms {
		args = append(args, platformFlag, platform)
	}

	return args
}

func getParentDirPath(name string) string {
	found, err := exec.LookPath(name)
	if err == nil {
		return path.Dir(found)
	}

	local, err := os.Getwd()
	if err == nil {
		return local
	}

	return defaultLocalPath
}
