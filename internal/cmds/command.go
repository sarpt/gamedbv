package cmds

import (
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
)

type command struct {
	name string
	args []string
	out  io.Writer
	err  io.Writer
	cmd  *exec.Cmd
}

func newCommand(name string, path string, args []string, out io.Writer, err io.Writer) command {
	if path == "" {
		path = getParentDirPath(name)
	}

	return command{
		name: name,
		args: args,
		out:  out,
		err:  err,
		cmd:  exec.Command(filepath.Join(path, name), args...),
	}
}

func (c command) Start() error {
	stdout, err := c.cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("could not open cmd stdout: %s", err)
	}

	stderr, err := c.cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("could not open cmd stderr: %s", err)
	}

	go func() {
		written, err := io.Copy(c.out, stdout)
		if err != nil {
			fmt.Printf("err: %s\n", err) // todo: some kind of better logging
		}
		fmt.Printf("written %d\n", written)
	}()

	go func() {
		writtenerr, err := io.Copy(c.err, stderr)
		if err != nil {
			fmt.Printf("err: %s\n", err)
		}
		fmt.Printf("writtenerr %d\n", writtenerr)
	}()

	return c.cmd.Start()
}

func (c command) Wait() error {
	return c.cmd.Wait()
}

func (c command) Execute() error {
	err := c.Start()
	if err != nil {
		return err
	}

	err = c.Wait()
	return err
}
