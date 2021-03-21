package network

import (
	messages "MP2/message"
	"encoding/gob"
	"net"
)

// Encode sends a gob encoded Message object through a TCP connection.
func encode(conn net.Conn, msg messages.Message) error {
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(msg)
	return err
}

// Decode receives a gob encoded Message object from a TCP connection and decodes it.
func decode(conn net.Conn, msg *messages.Message) error {
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(msg)
	return err
}
