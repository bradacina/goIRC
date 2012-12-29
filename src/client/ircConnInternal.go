package client

// This is the place where all IRCConn internal struct modifications are done

// internaly changes the topic of a specific channel
func (this *IRCConn) changeChannelTopic(chanName string, topic string) {

	found := false

	for _, val := range this.Channels {
		if val.Name == chanName {
			val.Topic = topic
			found = true
		}
	}

	if !found {
		log.Println("ERROR: Cannot find channel for which to change topic")
	}

}

// adds a channel to the internal IRCConn.Channels structure
func (this *IRCConn) addChannel(chanName string, topic string) {

	for _, val := range this.Channels {
		if val.Name == chanName {
			log.Println("ERROR: You are already on that channel")
			return
		}
	}

	channel := NewChannel(chanName)
	channel.Topic = topic

	this.Channels = append(this.Channels, channel)
}
