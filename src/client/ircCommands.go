package client

import "log"

// sends a message to the irc server
func (ircCon *IRCConn) SendMessage(message string) error {

	if message[len(message)-1] != '\n' {
		message = message + "\n"
	}

	log.Print(">>>", message)
	_, err := ircCon.con.Write([]byte(message))

	return err
}

// sends the NICK command to the irc server requesting a nickname change
func (ircCon *IRCConn) SetNickname(nickname string) error {

	msg := NICK + SPACE + nickname
	return ircCon.SendMessage(msg)
}

// sends the USER command to the irc server
func (ircCon *IRCConn) SetUser(username string) error {
	msg := USER + SPACE + username + SPACE + "3 * :" + username
	return ircCon.SendMessage(msg)
}

// set ourselves as invisible
func (ircCon *IRCConn) SetInvisible(nickname string) error {
	msg := MODE + SPACE + nickname + SPACE + "+i"
	return ircCon.SendMessage(msg)
}

// set ourselves as visible
func (ircCon *IRCConn) UnsetInvisible(nickname string) error {
	msg := MODE + SPACE + nickname + SPACE + "-i"
	return ircCon.SendMessage(msg)
}

// set ourselves as away
func (ircCon *IRCConn) SetAway(message string) error {
	if len(message) == 0 {
		message = ":Away"
	} else if message[0] != ':' {
		message = ":" + message
	}

	msg := AWAY + SPACE + message
	return ircCon.SendMessage(msg)
}

// unset ourselves from away
func (ircCon *IRCConn) UnsetAway() error {
	msg := AWAY
	return ircCon.SendMessage(msg)
}

func (ircCon *IRCConn) Quit(quitMsg string) error {
	msg := QUIT + SPACE + quitMsg
	return ircCon.SendMessage(msg)
}

// sends the JOIN command to the irc server requesting that we join a channel
func (ircCon *IRCConn) JoinChannel(channel string) error {
	msg := JOIN + SPACE + channel
	return ircCon.SendMessage(msg)
}

// sends the PART command to the irc server requesting that we part a channel
func (ircCon *IRCConn) PartChannel(channel, message string) error {
	if len(message) > 0 && message[0] != ':' {
		message = ":\"" + message + "\""
	}
	msg := PART + SPACE + channel + SPACE + message
	return ircCon.SendMessage(msg)
}

// sends the TOPIC command to the irc server requesting that we set a
// channels topic
func (ircCon *IRCConn) SetTopic(chanName string, topic string) error {

	var colon string

	if len(topic) > 0 && topic[0] != ':' {
		colon = ":"
	}

	msg := TOPIC + SPACE + chanName + SPACE + colon + topic
	return ircCon.SendMessage(msg)
}

// send the TOPIC command to the irc server requesting to retrieve a
// channels topic
func (ircCon *IRCConn) GetTopic(chanName string) error {
	msg := TOPIC + SPACE + chanName
	return ircCon.SendMessage(msg)
}
