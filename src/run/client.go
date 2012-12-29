package main

import "client"
import "fmt"

//import "time"
//import "log"

var replChan = make(chan string, 1)

var done = make(chan bool, 1)

var ircConn = client.NewIRCConn()

func main() {

	ircConn.Connect("irc.freenode.net:6667")

	go repl()

	go waitInput()

	<-done

	ircConn.Disconnect()
}

func repl() {
	for {
		select {
		case command := <-replChan:
			if command == "quit" {
				done <- true
			}

			if command == "j" {
				ircConn.JoinChannel("#animosity")
			}

			if command == "t" {
				ircConn.SetTopic("#animosity", "hi hi hi")
			}

		}
	}
}

func waitInput() {
	var input string
	for {
		n, err := fmt.Scanln(&input)
		if err == nil && n != 0 {
			replChan <- input
		}
	}
}
