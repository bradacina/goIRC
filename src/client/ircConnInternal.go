package client

// This is the place where all IRCConn internal struct modifications are done

import "log"

// internaly changes the topic of a specific channel
func (self *IRCConn) changeChannelTopic(chanName string, topic string) {

	val, ok := self.Channels[chanName]

	if !ok {
		log.Println("ERROR: Cannot find channel for which to change topic")
		return
	}

	val.Topic = topic
}

// adds a channel to the internal IRCConn.Channels structure
func (self *IRCConn) addChannel(chanName string, topic string) {

	_, ok := self.Channels[chanName]

	if ok {
		log.Println("ERROR: We have already added that channel")
		return
	}

	channel := NewChannel(chanName)
	channel.Topic = topic

	self.Channels[chanName] = channel
}

func (self *IRCConn) setChannelNames(chanName string, names []string) {

	val, ok := self.Channels[chanName]
	if !ok {
		log.Println("ERROR: Cannot find channel on which to set the names")
		return
	}

	log.Println("INTERNAL: Setting channel names for", chanName, "to", names)
	val.Names = names

}
