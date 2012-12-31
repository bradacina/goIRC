package client

import "net"
import "log"
import "strings"

type IRCConn struct {
	con               *net.TCPConn
	Nickname          string
	Username          string
	Channels          map[string]Channel
	AltNicknames      []string
	lastTriedNickname int
}

func NewIRCConn() *IRCConn {
	var c = new(IRCConn)

	c.Channels = make(map[string]Channel)

	// TODO: need to set these parameters in a config file or config step
	c.Username = "dssfds"

	c.AltNicknames = make([]string, 2)
	c.AltNicknames[0] = "Asdasd23244"
	c.AltNicknames[1] = "blahblah2324"

	return c
}

// connects to an irc server
// address if of the form "irc.freenode.net:6667"
func (self *IRCConn) Connect(address string) error {

	var addr, err = net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Println(err)
		return err
	}

	self.con, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println(err)
		return err
	}

	go self.listenForIncoming()

	self.SetNickname(self.AltNicknames[self.lastTriedNickname])

	self.SetUser(self.Username)

	return nil
}

func (self *IRCConn) Disconnect() error {
	return self.con.Close()
}

// listens for incoming messages from the irc server
func (self *IRCConn) listenForIncoming() {

	var b = make([]byte, 1024)
	var leftOver string
	var haveLeftOver bool

	for {
		n, err := self.con.Read(b)
		if err != nil {
			log.Println(err)
			self.Disconnect()
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
			self.translateMessage(val)
		}
	}
}

// tokenizes a message sent from the server, extracts the prefix, command
// and parameters and lets the handleCommand method do the rest
func (self *IRCConn) translateMessage(message string) {

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

	self.handleCommand(command, params)
}
