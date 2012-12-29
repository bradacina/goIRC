package client

type Channel struct {
	Name    string
	Topic   string
	Members []string
}

func NewChannel(name string) Channel {
	var c Channel

	c.Name = name

	return c
}
