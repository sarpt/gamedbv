package api

import (
	"encoding/json"
	"fmt"

	"github.com/sarpt/gamedbv/internal/progress"
)

const (
	progressState string = "progress"
	doneState     string = "done"
	errorState    string = "error"
)

type operationMessage struct {
	Op operation `json:"op"`
}

type clientOpertionMessage struct {
	operationMessage
	Payload interface{}
}

type startPayload struct {
	Platforms []string `json:"platforms"`
}

type startCommand struct {
	operationMessage
	startPayload
}

type closeCommand struct{}

func (clientCmd *clientOpertionMessage) UnmarshalJSON(data []byte) error {
	opValue := struct {
		Cmd *operation `json:"op"`
	}{}

	err := json.Unmarshal(data, &opValue)
	if err != nil {
		return err
	} else if opValue.Cmd == nil {
		return fmt.Errorf("message is not a operation instruction")
	}

	clientCmd.Op = *opValue.Cmd
	err = fillCommand(data, clientCmd)

	return err
}

func fillCommand(data []byte, clientOp *clientOpertionMessage) error {
	if clientOp.Op == startOp {
		cmd := startCommand{}

		err := json.Unmarshal(data, &cmd)
		if err != nil {
			return err
		}

		clientOp.Op = cmd.Op
		clientOp.Payload = startPayload{
			Platforms: cmd.Platforms,
		}

		return nil
	}

	return fmt.Errorf("operation '%s' not recognized", clientOp.Op)
}

type operationStatus struct {
	State string `json:"state"`
	progress.Status
}
