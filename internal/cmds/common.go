package cmds

import (
	"os"
	"os/exec"
	"path"
	"strconv"
)

const (
	defaultLocalPath = "."
	longArgPrefix    = "--"
)

func createPlatformsArguments(platforms []string) []string {
	return createMultipleFlagArguments(PlatformFlag, platforms)
}

func createTextArgument(text string) []string {
	return createSingleFlagArgument(TextFlag, text)
}

func createPageArgument(page int) []string {
	return createSingleFlagArgument(PageFlag, strconv.Itoa(page))
}

func createLimitArgument(limit int) []string {
	return createSingleFlagArgument(PageLimitFlag, strconv.Itoa(limit))
}

func createRegionsArguments(regions []string) []string {
	return createMultipleFlagArguments(RegionFlag, regions)
}

func createJSONArgument(isSet bool) []string {
	if isSet {
		return []string{longArgument(JSONFlag)}
	}

	return []string{}
}

func createMultipleFlagArguments(flagName string, values []string) []string {
	args := []string{}

	for _, val := range values {
		args = append(args, createSingleFlagArgument(flagName, val)...)
	}

	return args
}

func createSingleFlagArgument(flagName string, value string) []string {
	return []string{longArgument(flagName), value}
}

func longArgument(flagName string) string {
	return longArgPrefix + flagName
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
