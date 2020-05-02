package api

import (
	"sync"

	"github.com/gorilla/websocket"
)

type statusWriter struct {
	conn *websocket.Conn
	mu   *sync.Mutex
}

func newStatusWriter(conn *websocket.Conn) statusWriter {
	return statusWriter{
		conn: conn,
		mu:   &sync.Mutex{},
	}
}

func (sw statusWriter) Write(p []byte) (int, error) {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	status := statusMessage{
		Status:  "progress",
		Message: string(p),
	}
	err := sw.conn.WriteJSON(status)
	return len(p), err
}
