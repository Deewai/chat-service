FROM golang:latest

ARG SOCKET_PORT=8080
ENV SOCKET_PORT=$SOCKET_PORT

LABEL maintainer="Innocent Abdullahi <deewai48@gmail.com>"

RUN apt-get update


# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/Deewai/chat-service

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
ENV GO111MODULE=on
RUN go mod download
# RUN go mod vendor

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .
RUN go install -v ./...

# This container exposes port 8000 to the outside world
EXPOSE $SOCKET_PORT

# Run the executable
CMD ["chat-service"]
