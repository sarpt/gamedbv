package cmds

import (
	"io"
)

const idxName = "gamedbv-idx"

// Idx is used to execute dl component binary
type Idx struct {
	command command
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

// NewIdx returns new Idx cmd
func NewIdx(cfg IdxCfg, args IdxArguments) Idx {
	allArgs := createJSONArgument(true)
	allArgs = append(allArgs, createPlatformsArguments(args.Platforms)...)

	cmd := newCommand(idxName, cfg.Path, allArgs, cfg.Output, cfg.ErrOutput)
	return Idx{
		command: cmd,
	}
}

// Execute runs the command and waits for it to finish
func (idx Idx) Execute() error {
	return idx.command.Execute()
}
