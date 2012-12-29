package client

// sends a message to the irc server
func (this *IRCConn) SendMessage(message string) error {

	if message[len(message)-1] != '\n' {
		message = message + "\n"
	}

	log.Print(">>>", message)
	_, err := this.con.Write([]byte(message))

	return err
}

// sends the NICK command to the irc server requesting a nickname change
func (this *IRCConn) SetNickname(nickname string) error {

	msg := NICK + SPACE + nickname
	return this.SendMessage(msg)
}

// sends the USER command to the irc server
func (this *IRCConn) SetUser(username string) error {
	msg := USER + SPACE + username + SPACE + "3 * :" + username
	return this.SendMessage(msg)
}

// sends the JOIN command to the irc server requesting that we join a channel
func (this *IRCConn) JoinChannel(channel string) error {
	msg := JOIN + SPACE + channel
	return this.SendMessage(msg)
}

// sends the TOPIC command to the irc server requesting that we set a
// channels topic
func (this *IRCConn) SetTopic(chanName string, topic string) error {

	var colon string

	if len(topic) > 0 && topic[0] != ':' {
		colon = ":"
	}

	msg := TOPIC + SPACE + chanName + SPACE + colon + topic
	return this.SendMessage(msg)
}
