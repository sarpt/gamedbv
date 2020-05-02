package api

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type cmd string

const (
	startCmd cmd = "start"
	closeCmd cmd = "close"
)

var cmdHandlers = map[cmd]func(payload interface{}) error{
	startCmd: handleStartCmd,
}

func handleCmdMessage(msg clientCmdMessage) error {
	handler, ok := cmdHandlers[msg.Cmd]
	if !ok {
		return fmt.Errorf("no handler for the command: %s", msg.Cmd)
	}

	err := handler(msg.Payload)
	return err
}

func handleStartCmd(payload interface{}) error {
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

	// this should be handled with stdout and stderr separated (so one does not block the other), and with information being sent by websocket
	go func() {
		written, err := io.Copy(os.Stdout, out)
		if err != nil {
			fmt.Printf("err: %s\n", err)
		}
		fmt.Printf("written %d\n", written)

		writtenerr, err := io.Copy(os.Stdout, outerr)
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