package client

import (
	// "reflect"
	// "sync"
	// "fmt"
	"sync"
	"testing"

	// "time"

	// "time"

	// "github.com/gorilla/websocket"
	structures "github.com/Deewai/chat-service/structures"
	"github.com/stretchr/testify/assert"
	// "github.com/posener/wstest"
)

type myWebsocketConn struct {
	count int
}

func (w *myWebsocketConn) ReadMessage() (int, []byte, error) {
	if w.count == 0 {
		return 0, []byte("A test message"), nil
	}
	w.count++
	return 0, nil, nil
}
func (w *myWebsocketConn) WriteMessage(messageType int, data []byte) error {
	return nil
}
func (w *myWebsocketConn) WriteJSON(v interface{}) error {
	return nil
}
func (w *myWebsocketConn) Close() error {
	return nil
}

func TestNewConn(t *testing.T) {
	conn := NewConn(&myWebsocketConn{}, &Pool{}, "testID")
	assert.ObjectsAreEqual(&Instance{}, conn)
}

func TestInstance_Read(t *testing.T) {
	pool := &Pool{
		Broadcast: make(chan []byte),
		clients:   make(map[string]*Instance),
		Mutex:     sync.Mutex{},
		close:     make(chan interface{}),
	}
	instance1 := &Instance{
		ID:           "id123",
		conn:         &myWebsocketConn{},
		message:      make(chan []byte),
		LastMessages: structures.NewQueue(),
		Pool:         pool,
		close:        make(chan interface{}),
	}
	instance2 := &Instance{
		ID:           "id1234",
		conn:         &myWebsocketConn{},
		message:      make(chan []byte),
		LastMessages: structures.NewQueue(),
		Pool:         pool,
		close:        make(chan interface{}),
	}
	pool.NewClient(instance1)
	pool.NewClient(instance2)
	// go pool.StartClient(instance.ID)
	wg := sync.WaitGroup{}
	var gotten []byte
	go func() {
		for {
			select {
			case msg := <-instance1.message:
				gotten = msg
				wg.Done()
			}
		}
	}()
	message := []byte("A test message")
	wg.Add(1)
	instance2.processMessage(message)
	wg.Wait()
	instance1.closeConnection()
	instance2.closeConnection()
	pool = nil
	assert.EqualValues(t, message, gotten)
}
