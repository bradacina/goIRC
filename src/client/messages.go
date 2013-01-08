package client

// Define all handled or implemented messages to be sent to or received from
// the irc server

const PING = "PING"
const PONG = "PONG"
const NICK = "NICK"
const USER = "USER"
const JOIN = "JOIN"
const TOPIC = "TOPIC"
const MODE = "MODE"
const QUIT = "QUIT"
const AWAY = "AWAY"
const PART = "PART"

const ERR_NONICKNAMEGIVEN = "431"
const ERR_ERRONEUSNICKNAME = "432"
const ERR_NICKNAMEINUSE = "433"
const ERR_NICKCOLLISION = "436"
const ERR_UNAVAILRESOURCE = "437"

const RPL_NOTOPIC = "331"
const RPL_TOPIC = "332"
const RPL_NAMEREPLY = "353"

const SPACE = " "

func (ircCon *IRCConn) buildMessageMap() {

	ircCon.messageMap[PING] = (*IRCConn).handle_PING
	ircCon.messageMap[TOPIC] = (*IRCConn).handle_TOPIC
	ircCon.messageMap[RPL_TOPIC] = messageHandler((*IRCConn).handle_RPL_TOPIC)
	ircCon.messageMap[RPL_NOTOPIC] = messageHandler((*IRCConn).handle_RPL_TOPIC)
	ircCon.messageMap[RPL_NAMEREPLY] = messageHandler((*IRCConn).handle_RPL_NAMEREPLY)
	ircCon.messageMap[ERR_NONICKNAMEGIVEN] = messageHandler((*IRCConn).handle_NeedNewNickname)
	ircCon.messageMap[ERR_ERRONEUSNICKNAME] = messageHandler((*IRCConn).handle_NeedNewNickname)
	ircCon.messageMap[ERR_NICKNAMEINUSE] = messageHandler((*IRCConn).handle_NeedNewNickname)
	ircCon.messageMap[ERR_NICKCOLLISION] = messageHandler((*IRCConn).handle_NeedNewNickname)
	ircCon.messageMap[ERR_UNAVAILRESOURCE] = messageHandler((*IRCConn).handle_NeedNewNickname)
}
