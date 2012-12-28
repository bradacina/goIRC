package server

type InternalClientServerMessage struct {
	Client  *IRCClient
	Message string
}
