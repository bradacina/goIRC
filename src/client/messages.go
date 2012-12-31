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

const ERR_NONICKNAMEGIVEN = "431"
const ERR_ERRONEUSNICKNAME = "432"
const ERR_NICKNAMEINUSE = "433"
const ERR_NICKCOLLISION = "436"
const ERR_UNAVAILRESOURCE = "437"

const RPL_NOTOPIC = "331"
const RPL_TOPIC = "332"
const RPL_NAMEREPLY = "353"

const SPACE = " "
