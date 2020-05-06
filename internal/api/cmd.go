package api

import (
	"fmt"
	"io"
	"sync"

	"github.com/sarpt/gamedbv/internal/cmds"
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

	wg := sync.WaitGroup{}
	for _, platform := range startPayload.Platforms {
		wg.Add(1)

		go func(platform string) {
			defer wg.Done()

			err := updatePlatform(platform, w)
			if err != nil {
				fmt.Fprintf(w, "Update for platform %s failed", platform) // tbd: error writer
			}

			fmt.Fprintf(w, "Update for platform %s finished", platform)
		}(platform)
	}

	wg.Wait()

	return nil
}

func updatePlatform(platform string, w io.Writer) error {
	dlCfg := cmds.DlCfg{
		Output:    w,
		ErrOutput: w,
	}
	dlArgs := cmds.DlArguments{
		Platforms: []string{platform},
	}
	dlCmd := cmds.NewDl(dlCfg, dlArgs)

	err := dlCmd.Execute()
	if err != nil {
		return err
	}

	idxCfg := cmds.IdxCfg{
		Output:    w,
		ErrOutput: w,
	}
	idxArgs := cmds.IdxArguments{
		Platforms: []string{platform},
	}
	idxCmd := cmds.NewIdx(idxCfg, idxArgs)

	return idxCmd.Execute()
}
