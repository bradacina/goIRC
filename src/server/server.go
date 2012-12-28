// server project server.go
package server

import "net"
import "log"

type IRCServer struct {
	listenConn *net.TCPListener
	Done       chan bool
	Clients    []*IRCClient
	comChan    chan InternalClientServerMessage
}

func NewIRCServer() *IRCServer {

	s := new(IRCServer)
	s.Clients = make([]*IRCClient, 1)
	s.comChan = make(chan InternalClientServerMessage, 10)
	return s
}

func (this *IRCServer) Start() error {

	var addr, err = net.ResolveTCPAddr("tcp", "127.0.0.1:6667")
	if err != nil {
		return err
	}

	this.listenConn, err = net.ListenTCP("tcp", addr)

	if err != nil {
		return err
	}

	go this.accept()
	go this.processComChan()

	return nil
}

func (this *IRCServer) Stop() {
	this.listenConn.Close()
	this.Done <- true
}

func (this *IRCServer) accept() {
	for {
		con, err := this.listenConn.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}

		if c := NewIRCClient(con, this.comChan); c != nil {
			this.Clients = append(this.Clients, c)
		}
	}
}

func (this *IRCServer) processComChan() {
	for {
		select {
		case _ = <-this.comChan:

		}
	}
}
