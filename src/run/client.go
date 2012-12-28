package main

import "client"

func main() {
	ircConn := client.NewIRCConn()

	ircConn.Connect("irc.freenode.net:6667")

	wait := make(chan bool, 1)

	<-wait
}
