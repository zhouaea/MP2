package application

import (
	messages "MP2/message"
	"bufio"
	"fmt"
	"os"
)

// ReadClientCommandLineArgs parses the client initial command line arguments.
func ReadClientCommandLineArgs() (chatroomPort, username string) {
	// Read command line for id of process.
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Incorrect Usage. Do: go run clientprocess.go <chatroom_port> <username>\n")
		os.Exit(1)
	}

	if os.Args[2] == "chatroom" {
		fmt.Fprintf(os.Stderr, "'chatroom' is a reserved username. Please input another username\n")
		os.Exit(1)
	}

	// Parse process id from commandline.
	chatroomPort, username = os.Args[1], os.Args[2]
	return
}

// ReadClientCommandLineMessages reads the command line for messages to be routed by the chatroom.
func ReadClientCommandLineMessages(scanner *bufio.Scanner, username string) messages.Message {
	to, content := messagePrompt(scanner)
	message := messages.Message{to, username, content}
	return message
}

// messagePrompt prompts a user for a To, From, Title, and Content strings through the command line to create an email object.
func messagePrompt(scanner *bufio.Scanner) (to string, content string) {

	to = promptUser(scanner, "To: ")
	content = promptUser(scanner, "Content: ")

	return
}

// promptUser prints a prompt on a command line and reads a response.
// It will exit the program if the user types in 'EXIT'.
func promptUser(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	input := scanner.Text()
	if input == "EXIT" {
		fmt.Println("Exiting program per user input.")
		os.Exit(0)
	}
	return input
}