# chat-service
<!-- Author: Innocent Abdullahi <deewai48@gmail.com> -->

## Author: Innocent Abdullahi <deewai48@gmail.com>, github: https://hithub.com/Deewai

This is a prototype of a chat service implemented using websockets in Go

It makes use of the power of concurrency in Go, using buffered channels(for non-blocking code) for sending messages to different clients.

The chat is currently implemented to only send message received from a client to all other clients in a pool.

*Note*: This prototype does not include implementation of external services like databases

## Running Application
Application can be run with docker and docker-compose by simply running the below
```
SOCKET_PORT={port_to_listen_on} docker-compose up
```
The command above defaults to port 8080

The application can also be run without docker by executing the following commands
```
go build
SOCKET_PORT={port_to_listen_on} ./chat-service 
```

When a client connects he immediately receives a UUID generated ID that identifies him (No use case for the ID in this prototype tho)




