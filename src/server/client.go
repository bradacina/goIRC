package server

import "net"
import "log"
import "strings"

type IRCClient struct {
	con      *net.TCPConn
	comChan  chan InternalClientServerMessage
	Nickname string
}

func NewIRCClient(con *net.TCPConn, comChan chan InternalClientServerMessage) *IRCClient {

	c := new(IRCClient)
	c.con = con
	c.comChan = comChan

	go c.listenForData()
	return c
}

func (this *IRCClient) listenForData() {

	incoming := make([]byte, 1024)

	for {
		n, err := this.con.Read(incoming)
		if err != nil {
			log.Println(err)

			break
		}

		messages := strings.Split(string(incoming[:n]), "\n")

		for _, val := range messages {
			if len(val) == 0 {
				continue
			}
			internalMsg := InternalClientServerMessage{this, val}
			this.comChan <- internalMsg
		}

	}
}
