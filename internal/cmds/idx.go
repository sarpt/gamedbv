package cmds

import (
	"io"
)

const idxName = "gamedbv-idx"

// Idx is used to execute dl component binary
type Idx struct {
	command   command
	platforms []string
}

// IdxCfg is used to control the behavior of command executing the Dl component binary
type IdxCfg struct {
	Path      string
	Output    io.Writer
	ErrOutput io.Writer
}

// IdxArguments is used to provide arguments for command executing the Dl component binary
type IdxArguments struct {
	Platforms []string
}

// NewIdx returns
func NewIdx(cfg IdxCfg, args IdxArguments) Dl {
	allArgs := []string{}
	allArgs = append(allArgs, createPlatformArguments(args.Platforms)...)

	path := cfg.Path
	if path == "" {
		path = getParentDirPath(idxName)
	}

	cmd := newCommand(idxName, path, allArgs, cfg.Output, cfg.ErrOutput)
	return Dl{
		command:   cmd,
		platforms: args.Platforms,
	}
}

// Execute runs the command and waits for it to finish
func (idx Idx) Execute() error {
	return idx.command.Execute()
}
