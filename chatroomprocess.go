package main

import (
	"MP2/application"
	"MP2/network"
)

func main() {
	// Parse initial command line arguments.
	chatroomPort := application.ReadChatroomCommandLineArgs()

	// Listen to specified TCP port and handle client requests
	go network.ListenChatroomTCP(chatroomPort)

	// Wait for an exit message inputted through the command line.
	application.ReadChatroomCommand()
}