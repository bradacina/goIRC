package main

import "client"
import "fmt"

//import "time"
import "log"

var replChan = make(chan string, 1)

var done = make(chan bool, 1)

var ircConn = client.NewIRCConn()

var nickname = "asdas123456"

func main() {

	go listenCon()

	ircConn.Connect("irc.freenode.net:6667", nickname, "asd")

	go repl()

	go waitInput()

	<-done

	ircConn.Disconnect()
}

func listenCon() {
	for {
		select {
		case ntf := <-ircConn.ChannelNames():
			log.Println("MAIN: ", ntf.ChannelName, " ", ntf.Names)
		case ntf := <-ircConn.Topic():
			log.Println("MAIN: ", ntf.ChannelName, " ", ntf.Topic)
		case <-ircConn.NeedNickname():
			log.Println("MAIN: NEED NEW NICKNAME")
		}
	}
}

func repl() {
	for {
		select {
		case command := <-replChan:

			log.Println("INTERNAL: got user command:", command)
			if command == "quit" {
				done <- true
			}

			if command == "j" {
				ircConn.JoinChannel("#animosity")
			}

			if command == "t" {
				ircConn.SetTopic("#animosity", "hi hi hi")
			}

			if command == "-i" {
				ircConn.UnsetInvisible(nickname)
			}

			if command == "+i" {
				ircConn.SetInvisible(nickname)
			}

			if command == "q" {
				ircConn.Quit("")
			}

			if command == "A" {
				ircConn.SetAway("")
			}

			if command == "a" {
				ircConn.UnsetAway()
			}

			if command == "p" {
				ircConn.PartChannel("#animosity", "asd a sad asd")
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
