package client

// This is the place where we are handling all kinds of messages sent
// to us by the irc server

import "log"
import "errors"

// method to dispatch commands sent to us by the irc server
func (self *IRCConn) handleCommand(command string, params []string) {

	if command == PING {
		self.handle_PING(params)
	}

	if command == ERR_NICKNAMEINUSE ||
		command == ERR_NONICKNAMEGIVEN ||
		command == ERR_ERRONEUSNICKNAME ||
		command == ERR_NICKCOLLISION ||
		command == ERR_UNAVAILRESOURCE {

		self.handle_NeedNewNickname()
	}

	if command == RPL_TOPIC {
		self.handle_RPL_TOPIC(params)
	}

	if command == TOPIC {
		self.handle_TOPIC(params)
	}

	if command == RPL_NAMEREPLY {
		self.handle_RPL_NAMEREPLY(params)
	}
}

func (self *IRCConn) handle_RPL_NAMEREPLY(params []string) error {

	chanName, names, err := parse_RPL_NAMEREPLY(params)

	if err != nil {
		return err
	}

	self.addChannel(chanName, "")

	self.setChannelNames(chanName, names)

	return nil
}

func (self *IRCConn) handle_NeedNewNickname() {
	self.Nickname = ""

	self.lastTriedNickname += 1
	self.lastTriedNickname %= len(self.AltNicknames)

	self.SetNickname(self.AltNicknames[self.lastTriedNickname])
}

// method that handles the PING command sent from the server
func (self *IRCConn) handle_PING(params []string) error {

	if len(params) < 1 {
		log.Println("Error: Received PING with no parameters?")
		return errors.New("Error: Received PING with no parameters?")
	}

	msg := PONG + SPACE + params[0]
	return self.SendMessage(msg)
}

// method that handles the RPL_TOPIC command sent from the server
func (self *IRCConn) handle_RPL_TOPIC(params []string) error {
	chanName, topic, err := parse_RPL_TOPIC(params)

	if err != nil {
		return err
	}

	log.Println("Internal: Setting topic for", chanName, "to", topic)

	self.addChannel(chanName, topic)

	return nil
}

// method that handles the TOPIC command sent from the server
func (self *IRCConn) handle_TOPIC(params []string) error {
	chanName, topic, err := parse_TOPIC(params)

	if err != nil {
		return err
	}

	log.Println("Internal: Changing topic for", chanName, "to", topic)

	self.changeChannelTopic(chanName, topic)

	return nil
}
