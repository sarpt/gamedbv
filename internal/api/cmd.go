package api

import (
	"fmt"
	"io"
	"os/exec"
)

type cmd string

const (
	startCmd cmd = "start"
	closeCmd cmd = "close"
)

var cmdHandlers = map[cmd]func(payload interface{}, w io.Writer) error{
	startCmd: handleStartCmd,
}

func handleCmdMessage(msg clientCmdMessage, w io.Writer) error {
	handler, ok := cmdHandlers[msg.Cmd]
	if !ok {
		return fmt.Errorf("no handler for the '%s' command", msg.Cmd)
	}

	err := handler(msg.Payload, w)
	return err
}

func handleStartCmd(payload interface{}, w io.Writer) error {
	startPayload, ok := payload.(startPayload)
	if !ok || len(startPayload.Platforms) < 1 {
		return fmt.Errorf("incorrect payload for start command")
	}

	plat := startPayload.Platforms[0] // handle more than one platform (parallel?)
	cmd := exec.Command("./gamedbv-dl", "-platform", plat)

	out, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("could not open cmd stdout: %s", err)
	}

	outerr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("could not open cmd stderr: %s", err)
	}

	go func() {
		written, err := io.Copy(w, out)
		if err != nil {
			fmt.Printf("err: %s\n", err) // todo: some kind of better logging
		}
		fmt.Printf("written %d\n", written)
	}()

	go func() {
		writtenerr, err := io.Copy(w, outerr)
		if err != nil {
			fmt.Printf("err: %s\n", err)
		}
		fmt.Printf("writtenerr %d\n", writtenerr)
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
