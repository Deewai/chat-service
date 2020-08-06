package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Deewai/chat-service/client"
	"github.com/Deewai/chat-service/websocket_wrapper"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket_wrapper.Upwrapper{Upgrader: &websocket.Upgrader{}}

// Create a new pool of clients
var pool = client.NewPool()
var port = ":8080"

func main() {
	go pool.Start()
	router := mux.NewRouter()
	// Handle function /:id where id represents client's identification
	router.HandleFunc("/", handleSocketConnections)
	if p := os.Getenv("SOCKET_PORT"); p != "" {
		port = fmt.Sprintf(":%s", p)
	}
	log.Println("http server started on ", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleSocketConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		handleWebsocketError(err)
	}
	defer ws.Close()
	// send a generated clientID to the user
	clientID := uuid.New().String()
	ws.WriteJSON(map[string]string{"id": clientID})
	// Get a new instance of a client
	instance := client.NewConn(ws, pool, clientID)
	// Add client to client pool
	pool.NewClient(instance)
	// Start the client
	pool.StartClient(instance.ID)

}

func handleWebsocketError(err error) {
	if e, ok := err.(*websocket.CloseError); ok && (e.Code == websocket.CloseNormalClosure || e.Code == websocket.CloseNoStatusReceived) {
		log.Println("Websocket connection closed")
	}

}
