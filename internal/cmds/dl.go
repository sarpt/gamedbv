package cmds

import (
	"io"
)

const dlName = "gamedbv-dl"

// Dl is used to execute dl component binary
type Dl struct {
	command   command
	platforms []string
}

// DlCfg is used to control the behavior of command executing the Dl component binary
type DlCfg struct {
	Path      string
	Output    io.Writer
	ErrOutput io.Writer
}

// DlArguments is used to provide arguments for command executing the Dl component binary
type DlArguments struct {
	Platforms []string
}

// NewDl returns Dl command that is ready to be executed
func NewDl(cfg DlCfg, args DlArguments) Dl {
	allArgs := []string{jsonFlag}
	allArgs = append(allArgs, createPlatformArguments(args.Platforms)...)

	path := cfg.Path
	if path == "" {
		path = getParentDirPath(dlName)
	}

	cmd := newCommand(dlName, path, allArgs, cfg.Output, cfg.ErrOutput)
	return Dl{
		command:   cmd,
		platforms: args.Platforms,
	}
}

// Execute runs the command and waits for it to finish
func (dl Dl) Execute() error {
	return dl.command.Execute()
}
