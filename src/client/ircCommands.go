package client

import "log"

// sends a message to the irc server
func (self *IRCConn) SendMessage(message string) error {

	if message[len(message)-1] != '\n' {
		message = message + "\n"
	}

	log.Print(">>>", message)
	_, err := self.con.Write([]byte(message))

	return err
}

// sends the NICK command to the irc server requesting a nickname change
func (self *IRCConn) SetNickname(nickname string) error {

	msg := NICK + SPACE + nickname
	self.Nickname = nickname
	return self.SendMessage(msg)
}

// sends the USER command to the irc server
func (self *IRCConn) SetUser(username string) error {
	msg := USER + SPACE + username + SPACE + "3 * :" + username
	return self.SendMessage(msg)
}

// set ourselves as invisible
func (self *IRCConn) SetInvisible() error {
	msg := MODE + SPACE + self.Nickname + SPACE + "+i"
	return self.SendMessage(msg)
}

func (self *IRCConn) UnsetInvisible() error {
	msg := MODE + SPACE + self.Nickname + SPACE + "-i"
	return self.SendMessage(msg)
}

func (self *IRCConn) Quit(quitMsg string) error {
	msg := QUIT + SPACE + quitMsg
	return self.SendMessage(msg)
}

// sends the JOIN command to the irc server requesting that we join a channel
func (self *IRCConn) JoinChannel(channel string) error {
	msg := JOIN + SPACE + channel
	return self.SendMessage(msg)
}

// sends the TOPIC command to the irc server requesting that we set a
// channels topic
func (self *IRCConn) SetTopic(chanName string, topic string) error {

	var colon string

	if len(topic) > 0 && topic[0] != ':' {
		colon = ":"
	}

	msg := TOPIC + SPACE + chanName + SPACE + colon + topic
	return self.SendMessage(msg)
}

// send the TOPIC command to the irc server requesting to retrieve a
// channels topic
func (self *IRCConn) GetTopic(chanName string) error {
	msg := TOPIC + SPACE + chanName
	return self.SendMessage(msg)
}
