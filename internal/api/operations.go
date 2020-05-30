package api

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/sarpt/gamedbv/internal/cmds"
)

type operation string

const (
	startOp operation = "start"
	closeOp operation = "close"
)

var operationHandlers = map[operation]func(payload interface{}, w io.Writer) error{
	startOp: handleStartOperation,
}

func handleOperationMessage(msg clientOpertionMessage, w io.Writer) error {
	handler, ok := operationHandlers[msg.Op]
	if !ok {
		return fmt.Errorf("no handler for the '%s' operation", msg.Op)
	}

	err := handler(msg.Payload, w)
	return err
}

func handleStartOperation(payload interface{}, w io.Writer) error {
	startPayload, ok := payload.(startPayload)
	if !ok || len(startPayload.Platforms) < 1 {
		return fmt.Errorf("incorrect payload for start operation")
	}

	wg := sync.WaitGroup{}
	for _, platform := range startPayload.Platforms {
		wg.Add(1)

		go func(platform string) {
			defer wg.Done()

			err := updatePlatform(platform, w)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Update for platform %s failed: %v", platform, err) // tbd: tee writer
				fmt.Fprintf(w, "Update for platform %s failed", platform)                  // tbd: error writer
				return
			}

			status, err := json.Marshal(PlatformUpdateEndStatus(platform))
			if err != nil {
				fmt.Fprintf(w, "Error writing done status for platform %s", platform) // tbd: error writer
				return
			}

			fmt.Fprintf(w, string(status))
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
