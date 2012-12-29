package client

// This is the place where we are handling all kinds of messages sent
// to us by the irc server

// method to dispatch commands sent to us by the irc server
func (this *IRCConn) handleCommand(command string, params []string) {

	if command == PING {
		this.handle_PING(params)
	}

	if command == ERR_NICKNAMEINUSE {
		this.SetNickname("o435323234")
	}

	if command == RPL_TOPIC {
		this.handle_RPL_TOPIC(params)
	}

	if command == TOPIC {
		this.handle_TOPIC(params)
	}
}

// method that handles the PING command sent from the server
func (this *IRCConn) handle_PING(params []string) error {

	if len(params) < 1 {
		log.Println("Error: Received PING with no parameters?")
		return errors.New("Error: Received PING with no parameters?")
	}

	msg := PONG + SPACE + params[0]
	return this.SendMessage(msg)
}

// method that handles the RPL_TOPIC command sent from the server
func (this *IRCConn) handle_RPL_TOPIC(params []string) error {
	chanName, topic, err := parse_RPL_TOPIC(params)

	if err != nil {
		return err
	}

	log.Println("Setting topic for ", chanName, " to ", topic)

	this.addChannel(chanName, topic)

	return nil
}

// method that handles the TOPIC command sent from the server
func (this *IRCConn) handle_TOPIC(params []string) error {
	chanName, topic, err := parse_TOPIC(params)

	if err != nil {
		return err
	}

	log.Println("Changing topic for ", chanName, " to ", topic)

	this.changeChannelTopic(chanName, topic)

	return nil
}
