package client

import (
	"fmt"
	"sync"
	"testing"

	structures "github.com/Deewai/chat-service/structures"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestNewPool(t *testing.T) {
	pool := NewPool()
	assert.NotNil(t, pool)
}

func TestPool_NewClientSuccessfull(t *testing.T) {
	pool := &Pool{
		Broadcast: make(chan []byte),
		clients:   map[string]*Instance{},
		Mutex:     sync.Mutex{},
	}
	client := &Instance{ID: "", message: make(chan []byte), close: make(chan interface{}), Pool: pool}
	err := pool.NewClient(client)
	assert.NotNil(t, pool)
	assert.Len(t, pool.clients, 1)
	assert.Nil(t, err)
}

func TestPool_NewClientAlreadyExists(t *testing.T) {
	t.Run("Adding a duplicate client should throw an error of type *ClientError{}", func(t *testing.T) {
		pool := &Pool{
			Broadcast: make(chan []byte),
			clients:   map[string]*Instance{},
			Mutex:     sync.Mutex{},
		}
		id := "testID"
		pool.NewClient(&Instance{ID: id})
		err := pool.NewClient(&Instance{ID: id})
		assert.NotNil(t, err)
		assert.IsType(t, &ClientError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("An error occured with client id %s: %s", id, "Client already exists"))
	})

}

func TestPool_DeleteClientSuccessfull(t *testing.T) {
	t.Run("Adding a duplicate client should throw an error of type *ClientError{}", func(t *testing.T) {
		pool := &Pool{
			Broadcast: make(chan []byte),
			clients:   map[string]*Instance{},
			Mutex:     sync.Mutex{},
		}
		id := "testID"
		pool.NewClient(&Instance{ID: id})
		assert.Len(t, pool.clients, 1)
		err := pool.DeleteClient(id)
		assert.Nil(t, err)
		assert.Len(t, pool.clients, 0)
	})
}

func TestPool_DeleteClientNotFound(t *testing.T) {
	t.Run("Adding a duplicate client should throw an error of type *ClientError{}", func(t *testing.T) {
		pool := &Pool{
			Broadcast: make(chan []byte),
			clients:   map[string]*Instance{},
			Mutex:     sync.Mutex{},
		}
		id := "testID"
		err := pool.DeleteClient(id)
		assert.NotNil(t, err)
		assert.IsType(t, &ClientError{}, err)
		assert.EqualError(t, err, fmt.Sprintf("An error occured with client id %s: %s", id, "Client not found"))
	})

}

func TestPool_Start(t *testing.T) {
	pool := &Pool{
		Broadcast: make(chan []byte),
		clients:   map[string]*Instance{},
		Mutex:     sync.Mutex{},
	}
	client := &Instance{
		ID:           "test123",
		conn:         &websocket.Conn{},
		message:      make(chan []byte),
		Mutex:        sync.Mutex{},
		LastMessages: structures.NewQueue(),
	}
	pool.NewClient(client)
	go pool.Start()
	newMessage := []byte("a new message")
	pool.Broadcast <- newMessage
	assert.EqualValues(t, newMessage, <-client.message)
}
