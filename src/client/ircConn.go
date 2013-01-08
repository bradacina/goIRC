package client

import "net"
import "log"
import "strings"

type messageHandler func(*IRCConn, string, []string) error

type IRCConn struct {
	con *net.TCPConn

	channelNames chan ChannelNamesNotification
	topic        chan TopicNotification
	needNickname chan bool

	// keeps a map between incoming messages and their handler functions
	messageMap map[string]messageHandler
}

func NewIRCConn() *IRCConn {
	var c = new(IRCConn)

	c.messageMap = make(map[string]messageHandler)

	c.buildMessageMap()

	c.channelNames = make(chan ChannelNamesNotification, 10)
	c.topic = make(chan TopicNotification, 10)
	c.needNickname = make(chan bool, 10)

	return c
}

// returns a chan on which we send notifications when channel names are received
// from the irc server
func (ircCon *IRCConn) ChannelNames() <-chan ChannelNamesNotification {
	return ircCon.channelNames
}

// returns a chan on which we send notifications when channel topic changes are
// received from the irc server
func (ircCon *IRCConn) Topic() <-chan TopicNotification {
	return ircCon.topic
}

// returns a chan on which we send notifications when our nickname needs to be changed
func (ircCon *IRCConn) NeedNickname() <-chan bool {
	return ircCon.needNickname
}

// connects to an irc server
// address if of the form "irc.freenode.net:6667"
func (ircCon *IRCConn) Connect(address string, nickname string, user string) error {

	var addr, err = net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Println(err)
		return err
	}

	ircCon.con, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println(err)
		return err
	}

	go ircCon.listenForIncoming()

	ircCon.SetNickname(nickname)

	ircCon.SetUser(user)

	return nil
}

func (ircCon *IRCConn) Disconnect() error {
	return ircCon.con.Close()
}

// listens for incoming messages from the irc server
func (ircCon *IRCConn) listenForIncoming() {

	var b = make([]byte, 1024)
	var leftOver string
	var haveLeftOver bool

	for {
		n, err := ircCon.con.Read(b)
		if err != nil {
			log.Println(err)
			ircCon.Disconnect()
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
			ircCon.translateMessage(val)
		}
	}
}

// tokenizes a message sent from the server, extracts the prefix, command
// and parameters and lets the handleCommand method do the rest
func (ircCon *IRCConn) translateMessage(message string) {

	initialTokens := strings.Split(message, " ")

	var tokens []string

	// get rid of empty tokens
	for _, val := range initialTokens {
		if len(val) != 0 {
			tokens = append(tokens, val)
		}
	}

	var command string
	var prefix string
	var params []string

	if len(tokens) >= 1 {
		if (tokens[0])[0] == ':' {
			prefix = tokens[0]

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

	ircCon.handleCommand(prefix, command, params)
}
