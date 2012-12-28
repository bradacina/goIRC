// client project client.go
package client

import "net"
import "log"
import "strings"

//import "time"

type IRCConn struct {
	con      *net.TCPConn
	Nickname string
	Username string
}

func NewIRCConn() *IRCConn {
	var c = new(IRCConn)
	c.Nickname = "Asdasd23444"
	c.Username = "dssfds"

	return c
}

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

func (this *IRCConn) SendMessage(message string) error {

	if message[len(message)-1] != '\n' {
		message = message + "\n"
	}

	log.Print(">>>", message)
	_, err := this.con.Write([]byte(message))

	return err
}

func (this *IRCConn) SetNickname(nickname string) error {

	msg := "NICK " + nickname
	return this.SendMessage(msg)
}

func (this *IRCConn) SetUser(username string) error {
	msg := "USER " + username + " " + "3 * :" + username
	return this.SendMessage(msg)
}

func (this *IRCConn) handlePing(params string) error {
	msg := "PONG " + params
	return this.SendMessage(msg)
}

func (this *IRCConn) translateMessage(message string) {
	initialTokens := strings.Split(message, " ")

	var tokens []string

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
				// bad message format
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

	if command == "PING" {
		this.handlePing(params[0])
	}

	if command == "433" {
		this.SetNickname("o435323234")
	}
}
