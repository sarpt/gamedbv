package api

import (
	"sync"

	"github.com/gorilla/websocket"
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

func (sw progressWriter) Write(p []byte) (int, error) {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	status := statusMessage{
		Status:  "progress",
		Message: string(p),
	}
	err := sw.conn.WriteJSON(status)
	return len(p), err
}
