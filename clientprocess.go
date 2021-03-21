package main

import (
	"MP2/application"
	"MP2/network"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Parse command line arguments.
	chatroomPort, username := application.ReadClientCommandLineArgs()

	// Connect to chatroom server via TCP and send process username.
	channel := network.ConnectToChatroom(chatroomPort, username)
	
	// Listen to chatroom server messages. Close the process if the server declares that it is closed.
	go network.ListenToChatroomServer(channel, username)

	fmt.Println("Please input each message field on a separate line down below. To exit the program, type EXIT.")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Read commands from user via command line. Exit process if user types in 'EXIT'.
		message := application.ReadClientCommandLineMessages(scanner, username)

		// Send message through commandline.
		go network.SendToChatroom(channel, message)
	}
}