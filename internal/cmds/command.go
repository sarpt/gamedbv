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
	cmd  *exec.Cmd
}

func newCommand(name string, path string, args []string) command {
	if path == "" {
		path = getParentDirPath(name)
	}

	return command{
		name: name,
		args: args,
		cmd:  exec.Command(filepath.Join(path, name), args...),
	}
}

func (c command) InitializeWriters(outWriter io.Writer, errWriter io.Writer) error {
	stdout, err := c.cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("could not open cmd stdout: %s", err)
	}

	stderr, err := c.cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("could not open cmd stderr: %s", err)
	}

	go func() {
		written, err := io.Copy(outWriter, stdout)
		if err != nil {
			fmt.Printf("err: %s\n", err) // todo: some kind of better logging
		}
		fmt.Printf("written %d\n", written)
	}()

	go func() {
		writtenerr, err := io.Copy(errWriter, stderr)
		if err != nil {
			fmt.Printf("err: %s\n", err)
		}
		fmt.Printf("writtenerr %d\n", writtenerr)
	}()

	return nil
}

func (c command) Start() error {
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

func (c command) Stdout() ([]byte, error) {
	return c.cmd.Output()
}
