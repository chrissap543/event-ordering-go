package internal

import (
	"net"
)

type Connection struct {
	Name       string
	Connection net.Conn
}

type Message struct {
	Message string
}

func Multicast(connections []Connection, msg Message) {

}
