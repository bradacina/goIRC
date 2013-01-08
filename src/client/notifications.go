package client

type ChannelNamesNotification struct {
	ChannelName string
	Names       []string
}

type TopicNotification struct {
	ChannelName string
	Topic       string
}
