package client

// This is where we parse the parameters of a command sent to us by the irc
// server into meaningful information to be used by the handler methods

import "errors"
import "strings"

func parse_RPL_NAMEREPLY(params []string) (chanName string, names []string, err error) {
	if len(params) < 4 {
		return "", nil, errors.New("Not enough parameters")
	}

	chanName = params[2]

	for _, val := range params[3:] {
		if len(val) == 0 {
			continue
		}

		name := strings.Trim(val, "@+:\n\r")

		names = append(names, name)
	}

	err = nil
	return
}

func parse_RPL_TOPIC(params []string) (chanName string, topic string, err error) {
	if len(params) < 3 {
		return "", "", errors.New("Not enough parameters")
	}

	chanName = params[1]

	topic = strings.Trim(params[2], ":")

	for _, val := range params[3:] {
		topic = topic + SPACE + val
	}

	return
}

func parse_TOPIC(params []string) (chanName string, topic string, err error) {
	if len(params) < 2 {
		err = errors.New("Not enough parameters")
		return
	}

	chanName = params[0]

	topic = strings.Trim(params[1], ":")
	for _, val := range params[2:] {
		topic = topic + SPACE + val
	}

	return
}
