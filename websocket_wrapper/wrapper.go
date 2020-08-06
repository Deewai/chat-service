package websocket_wrapper

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Conn - Interface specifying methods needed to be satisfied by the websocket connection
type Conn interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	WriteJSON(v interface{}) error
	Close() error
}

// Upwrapper - Upgrader wrapper
type Upwrapper struct {
	Upgrader    *websocket.Upgrader
	CheckOrigin func(r *http.Request) bool
}

// Upgrade - expose the websocket upgrader function
func (a Upwrapper) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (Conn, error) {
	a.Upgrader.CheckOrigin = a.CheckOrigin
	return a.Upgrader.Upgrade(w, r, responseHeader)
}
