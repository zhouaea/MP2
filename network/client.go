package network

import (
	"MP2/errorchecker"
	messages "MP2/message"
	"fmt"
	"net"
	"os"
)

// ConnectToChatroom attempts to connect to a chatroom via a TCP port number and sends the client's username to be
// registered into the chatroom.
func ConnectToChatroom(chatroomPort string, username string) net.Conn {
	// Attempt to connect to the chatroom TCP channel on the localhost IP address.
	port := ":" + chatroomPort
	channel, err := net.Dial("tcp", port)
	errorchecker.CheckError(err)

	// Send chatroom username and port number.
	message := messages.Message{To: "chatroom", From: username}
	encode(channel, message)
	fmt.Printf("Username '%s' registered to chatroom at port %s\n", username, port)

	return channel
}

// ListenToChatroomServer reads and prints messages sent by the chatroom server.
// It will close the client process upon the chatroom's closing.
func ListenToChatroomServer(channel net.Conn, username string) {
	defer channel.Close()
	for {
		// Read and print message sent by chatroom server through TCP channel.
		message := new(messages.Message)
		err := decode(channel, message)

		// If there is an error reading from the server, we know the server is closed.
		if err != nil {
			fmt.Println("\nServer has closed.")
			os.Exit(0)
		}

		fmt.Println("\n---------------")
		fmt.Printf("From: %s\n", message.From)
		fmt.Println(message.Content)
		fmt.Println("---------------")
	}
}

// SendToChatroom sends a message from a client to a chatroom.
func SendToChatroom(channel net.Conn, message messages.Message) {
	encode(channel, message)
}