package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Command string

const (
	CMDSet Command = "SET"
	CMDGet Command = "GET"
	CMDDel Command = "DEL"
	CMDHas Command = "HAS"
)

type Message struct {
	Command Command
	Key     []byte
	Value   []byte
	TTL     time.Duration
}

func parseMessage(rawCmd []byte) (*Message, error) {
	strCmd := strings.Split(string(rawCmd), " ")
	if len(strCmd) < 2 {
		return nil, errors.New("invalid protocol format")
	}

	msg := &Message{
		Command: Command(strCmd[0]),
		Key:     []byte(strCmd[1]),
	}

	if msg.Command == CMDSet {
		if len(strCmd) < 4 {
			return nil, fmt.Errorf("need 4 arguments for SET command, got %d", len(strCmd))
		}
		ttl, err := strconv.Atoi(strCmd[3])
		if err != nil {
			return nil, errors.New("invalid TTL, must be a number")
		}

		msg.Value = []byte(strCmd[2])
		msg.TTL = time.Duration(ttl) * time.Millisecond
	}

	return msg, nil
}
