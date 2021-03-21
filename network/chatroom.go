package network

import (
	"MP2/errorchecker"
	messages "MP2/message"
	"fmt"
	"net"
)

// Stores client username to tcp connection pairs.
var clientLookup = make(map[string]net.Conn)
var repeat = 0

// ListenChatroomTCP configures a chatroom to listen for tcp connections and handles their messages,
func ListenChatroomTCP(chatroomPort string) {
	// Listen to an unused TCP port on localhost.
	port := ":" + chatroomPort
	listener, err := net.Listen("tcp", port)
	errorchecker.CheckError(err)
	fmt.Println("Listening to tcp port " + port + " was successful!")

	defer listener.Close()

	// Wait for a connection from a client and serve their connection until they disconnect.
	for {
		conn, err := listener.Accept()
		errorchecker.CheckError(err)

		go handleClientMessages(conn)
	}
}

// handleClientMessages does one of two actions after decoding a client message:
//	1. If message is addressed to chatroom, store a client's username and TCP connection as a key value pair.
//	2. If message is addressed to another client, send the message to the destination.
//	   If the recipient is not connected, send an error message to the sender.
func handleClientMessages(senderChannel net.Conn) {
	defer senderChannel.Close()
	// Keep the connection to clients going until a client leaves or the chatroom closes.
	for {
		// Decode message.
		message := new(messages.Message)
		err := decode(senderChannel, message)

		// If client disconnects, close connection and delete key-connection pair.
		if err != nil {
			break
		}

		// If message is addressed to chatroom, store a client's username and TCP connection as a key value pair.
		if message.To == "chatroom" {
			// The From field will have the username and the Content field will have the tcp port.
			clientLookup[message.From] = senderChannel
			fmt.Printf("Connection to %s was succesful!\n", message.From)

			// Delete username-connection pair when client disconnects.
			defer delete(clientLookup, message.From)
			defer fmt.Printf("%s has disconnected.\n", message.From)
			continue
		}

		// If message is addressed to another client, send the message to the destination.

		// If username-connection pair does not exist, do not attempt to send a message through a TCP connection, but
		// let the sender know there was an error
		if clientLookup[message.To] == nil {
			fmt.Println("Error connecting to message destination, no username connection pair")

			errorMessageContent := fmt.Sprintf("'%s' is not connected to the chatroom. Your message was not sent. \nFailed Message Contents:\n%s",
				message.To, message.Content)
			errorMessage := messages.Message{To: message.From, From: "chatroom", Content: errorMessageContent}
			encode(senderChannel, errorMessage)

			continue
		}

		// Attempt to send message to destination.
		err = encode(clientLookup[message.To], *message)

		// If an error somehow occurs when sending a message to a registered connection, let the sender know.
		if err != nil {
			fmt.Println("Error connecting to message destination, username exists but connection is not working")

			errorMessageContent := fmt.Sprintf("'%s' is not connected to the chatroom. Your message was not sent. \nFailed Message Contents:\n%s",
				message.To, message.Content)
			errorMessage := messages.Message{To: message.From, From: "chatroom", Content: errorMessageContent}
			encode(senderChannel, errorMessage)

			continue
		}
		fmt.Printf("Message sent from %s to %s.\n", message.From, message.To)
	}
}