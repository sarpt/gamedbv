package api

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/sarpt/gamedbv/internal/progress"
)

type progressWriter struct {
	conn *websocket.Conn
	mu   *sync.Mutex
}

func newProgressWriter(conn *websocket.Conn) progressWriter {
	return progressWriter{
		conn: conn,
		mu:   &sync.Mutex{},
	}
}

func (sw progressWriter) Write(payload []byte) (int, error) {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	if !json.Valid(payload) {
		return 0, fmt.Errorf("the progress payload is not a correct json")
	}

	var status progress.Status
	err := json.Unmarshal(payload, &status)
	if err != nil {
		return 0, err
	}

	statusMessage := statusMessage{
		State:  progressState,
		Status: status,
	}
	err = sw.conn.WriteJSON(statusMessage)
	return len(payload), err // todo: len(payload) is not correct, but the value is yet to be used. needs to be fixed
}
