package client

// This is the place where we are handling all kinds of messages sent
// to us by the irc server

import "log"
import "errors"

// method to dispatch commands sent to us by the irc server
func (ircCon *IRCConn) handleCommand(prefix, command string, params []string) {

	handler, b := ircCon.messageMap[command]
	if b {
		handler(ircCon, prefix, params)
	}

}

func (ircCon *IRCConn) handle_RPL_NAMEREPLY(prefix string, params []string) error {

	chanName, names, err := parse_RPL_NAMEREPLY(params)

	if err != nil {
		return err
	}

	ntf := ChannelNamesNotification{chanName, names}

	ircCon.channelNames <- ntf

	return nil
}

func (ircCon *IRCConn) handle_NeedNewNickname(prefix string, params []string) error {
	ircCon.needNickname <- true

	return nil
}

// method that handles the RPL_TOPIC command sent from the server
func (ircCon *IRCConn) handle_RPL_TOPIC(prefix string, params []string) error {
	chanName, topic, err := parse_RPL_TOPIC(params)

	if err != nil {
		return err
	}

	ntf := TopicNotification{chanName, topic}

	ircCon.topic <- ntf
	return nil
}

// method that handles the TOPIC command sent from the server
func (ircCon *IRCConn) handle_TOPIC(prefix string, params []string) error {
	chanName, topic, err := parse_TOPIC(params)

	if err != nil {
		return err
	}

	ntf := TopicNotification{chanName, topic}
	ircCon.topic <- ntf
	return nil
}

// method that handles the PING command sent from the server
func (ircCon *IRCConn) handle_PING(prefix string, params []string) error {

	if len(params) < 1 {
		ircCon.SendMessage(PONG)
		log.Println("Error: Received PING with no parameters?")
		return errors.New("Error: Received PING with no parameters?")
	}

	msg := PONG + SPACE + params[0]
	return ircCon.SendMessage(msg)
}
