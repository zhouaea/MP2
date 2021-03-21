package application

import (
	"bufio"
	"fmt"
	"os"
)

// ReadChatroomCommandLineArgs() parses the server command line arguments.
func ReadChatroomCommandLineArgs() (chatroomPort string) {
	// Read command line for id of process.
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Incorrect Usage. Do: go run chatroomprocess.go <chatroom_port>\n")
		os.Exit(1)
	}

	// Parse process id from commandline.
	chatroomPort = os.Args[1]
	return
}

// ReadChatroomCommand() waits for the user to type 'EXIT' in the chatroom command line, then terminating the chatroom.
func ReadChatroomCommand() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Type 'EXIT' to close the chatroom server.")
		scanner.Scan()
		input := scanner.Text()
		if input == "EXIT" {
			os.Exit(0)
		}
	}
}
