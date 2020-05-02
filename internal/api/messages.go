package api

import (
	"encoding/json"
	"fmt"
)

type cmdMessage struct {
	Cmd cmd `json:"cmd"`
}

type clientCmdMessage struct {
	cmdMessage
	Payload interface{}
}

type startPayload struct {
	Platforms []string `json:"platforms"`
}

type startCommand struct {
	cmdMessage
	startPayload
}

type closeCommand struct{}

func (clientCmd *clientCmdMessage) UnmarshalJSON(data []byte) error {
	cmdValue := struct {
		Cmd *cmd `json:"cmd"`
	}{}

	err := json.Unmarshal(data, &cmdValue)
	if err != nil {
		return err
	} else if cmdValue.Cmd == nil {
		return fmt.Errorf("command was not provided with the message")
	}

	clientCmd.Cmd = *cmdValue.Cmd
	err = fillCommand(data, clientCmd)

	return err
}

func fillCommand(data []byte, clientCmd *clientCmdMessage) error {
	if clientCmd.Cmd == startCmd {
		cmd := startCommand{}

		err := json.Unmarshal(data, &cmd)
		if err != nil {
			return err
		}

		clientCmd.Cmd = cmd.Cmd
		clientCmd.Payload = startPayload{
			Platforms: cmd.Platforms,
		}

		return nil
	}

	return fmt.Errorf("command not recognized: %s", clientCmd.Cmd)
}

type statusMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
