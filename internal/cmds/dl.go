package cmds

import (
	"io"
)

const dlName = "gamedbv-dl"

// Dl is used to execute dl component binary
type Dl struct {
	command   command
	output    io.Writer
	errOutput io.Writer
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
	allArgs := createJSONArgument(true)
	allArgs = append(allArgs, createPlatformsArguments(args.Platforms)...)

	cmd := newCommand(dlName, cfg.Path, allArgs)
	return Dl{
		command:   cmd,
		output:    cfg.Output,
		errOutput: cfg.ErrOutput,
	}
}

// Execute runs the command and waits for it to finish
func (dl Dl) Execute() error {
	err := dl.command.InitializeWriters(dl.output, dl.errOutput)
	if err != nil {
		return err
	}

	return dl.command.Execute()
}
