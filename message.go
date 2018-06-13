// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"github.com/intel-go/fastjson"
)

// MessageType const type of a message
type MessageType int

const (
	PingMessage MessageType = iota
	PingReqMessage
	JoinMessage
	LeftMessage
	FullMembersMessage
	GossipMessage
)

// Message is used for talking between members
type Message struct {
	Type MessageType `json:"type"`
	// reprent an ip address, PingReqMessage may use this to define ping target
	Target []byte `json:"target"`
	Nodes  []Node `json:"nodes,omitempty"`
}

// Valid check whether message is valid
func (message Message) Valid() bool {
	// message must have a type and should be in predefined type
	if !(message.Type >= PingMessage && message.Type <= GossipMessage) {
		return false
	}

	return true
}

func encode(message Message) (bytes []byte, err error) {
	return fastjson.Marshal(&message)
}

func decode(bytes []byte, message Message) (err error) {
	return fastjson.Unmarshal(bytes, &message)
}

func GetPingMessage() Message {
	return Message{
		Type: PingMessage,
	}
}
