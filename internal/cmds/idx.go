package cmds

import (
	"io"
)

const idxName = "gamedbv-idx"

// Idx is used to execute dl component binary
type Idx struct {
	command
	output    io.Writer
	errOutput io.Writer
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
	GRPC      bool
}

// NewIdx returns new Idx cmd
func NewIdx(cfg IdxCfg, args IdxArguments) Idx {
	allArgs := createJSONArgument(true)
	allArgs = append(allArgs, createGRPCFlag(args.GRPC)...)
	allArgs = append(allArgs, createPlatformsArguments(args.Platforms)...)

	command := newCommand(idxName, cfg.Path, allArgs)
	return Idx{
		command:   command,
		output:    cfg.Output,
		errOutput: cfg.ErrOutput,
	}
}

// Execute runs the command and waits for it to finish
func (idx Idx) Execute() error {
	err := idx.command.InitializeWriters(idx.output, idx.errOutput)
	if err != nil {
		return err
	}

	return idx.command.Execute()
}
