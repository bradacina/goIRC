package client

import "net"
import "log"
import "strings"
import "errors"

type IRCConn struct {
	con      *net.TCPConn
	Nickname string
	Username string
	Channels []Channel
}

func NewIRCConn() *IRCConn {
	var c = new(IRCConn)

	// TODO: need to set these parameters in a config file or config step
	c.Nickname = "Asdasd23444"
	c.Username = "dssfds"

	return c
}

// connects to an irc server
// address if of the form "irc.freenode.net:6667"
func (this *IRCConn) Connect(address string) error {

	var addr, err = net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Println(err)
		return err
	}

	this.con, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println(err)
		return err
	}

	go this.listenForIncoming()

	this.SetNickname(this.Nickname)

	this.SetUser(this.Username)

	return nil
}

func (this *IRCConn) Disconnect() error {
	return this.con.Close()
}

// listens for incoming messages from the irc server
func (this *IRCConn) listenForIncoming() {

	var b = make([]byte, 1024)
	var leftOver string
	var haveLeftOver bool

	for {
		n, err := this.con.Read(b)
		if err != nil {
			log.Println(err)
			this.Disconnect()
			break
		}

		bigMessage := string(b[:n])

		if haveLeftOver {
			bigMessage = leftOver + bigMessage
			haveLeftOver = false
		}

		if bigMessage[len(bigMessage)-1] != '\n' {
			haveLeftOver = true
		}

		messages := strings.Split(bigMessage, "\n")

		if haveLeftOver {
			leftOver = messages[len(messages)-1]
			messages = messages[:len(messages)-1]
		}

		for _, val := range messages {
			if len(val) == 0 {
				continue
			}
			log.Println("<<<", val)
			this.translateMessage(val)
		}
	}
}

// tokenizes a message sent from the server, extracts the prefix, command
// and parameters and lets the handleCommand method do the rest
func (this *IRCConn) translateMessage(message string) {

	initialTokens := strings.Split(message, " ")

	var tokens []string

	// get rid of empty tokens
	for _, val := range initialTokens {
		if len(val) != 0 {
			tokens = append(tokens, val)
		}
	}

	var command string
	//var prefix string
	var params []string

	if len(tokens) >= 1 {
		if (tokens[0])[0] == ':' {
			_ = tokens[0]

			if len(tokens) < 2 {
				log.Println("Error: Bad message format. Got Prefix but no Command.")
				return
			}

			command = tokens[1]

			if len(tokens) > 2 {
				params = tokens[2:]
			}

		} else {
			command = tokens[0]

			if len(tokens) > 1 {
				params = tokens[1:]
			}
		}
	}

	this.handleCommand(command, params)
}
