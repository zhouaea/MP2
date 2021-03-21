package messages

// When initializing username-connection pairs, if the To field is for "chatroom", then the chatroom process
//read the From field as the client username and ignores the Content field.
type Message struct {
	To      string
	From    string
	Content string
}
