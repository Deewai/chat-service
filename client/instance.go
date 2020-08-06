package client

import (
	"log"
	"sync"

	structures "github.com/Deewai/chat-service/structures"
	gwebsocket "github.com/gorilla/websocket"
)

const (
	messageBuffer = 10
)

type Conn interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	WriteJSON(v interface{}) error
	Close() error
}

// Instance type which is a client
type Instance struct {
	ID      string
	conn    Conn
	message chan []byte
	close   chan interface{}
	sync.Mutex
	LastMessages *structures.Queue
	Pool         pool
}

// Create a new connection/client
func NewConn(ws Conn, pool pool, ID string) *Instance {
	return &Instance{
		ID:           ID,
		conn:         ws,
		message:      make(chan []byte, messageBuffer),
		close:        make(chan interface{}),
		LastMessages: structures.NewQueue(),
		Pool:         pool,
	}
}

func (i *Instance) Read() {
	defer i.closeConnection()

	for {

		// Read in a new message byte array
		_, message, err := i.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			if _, ok := err.(*gwebsocket.CloseError); ok {
				// Close client SendMessage routine
				// close(i.close)

				// err = i.Pool.DeleteClient(i.ID)
				// if err != nil {
				// 	log.Println(err)
				// }
				break
			}
			// break
		} else {
			i.processMessage(message)
		}
	}
}

func (i *Instance) closeConnection() {
	err := i.conn.Close()
	if err != nil {
		log.Printf("Error closing instance connection: %v", err)
	}
	// Close client SendMessage routine
	close(i.close)
	// Delete client from pool
	err = i.Pool.DeleteClient(i.ID)
	if err != nil {
		log.Println(err)
	}

}

func (i *Instance) processMessage(message []byte) {
	// Send to all clients
	for _, client := range i.Pool.GetClients() {
		if client.ID != i.ID {
			client.message <- message
		}
	}
}

func (i *Instance) SendMessage() {
	for {
		select {
		case msg, _ := <-i.message:
			// Grab the next message from the message channel
			log.Println(msg)
			// Send it out to instance client that is currently connected
			i.conn.WriteMessage(gwebsocket.TextMessage, msg)
		case <-i.close:
			return
		}

	}
}
