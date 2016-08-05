// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"net"
)

// TCPTalk propagate with tcp protocol
type TCPTalk struct {
	conn net.Conn

	// heartbeat interval
	interval int64

	// timeout
	connectTimeout int64
	speakTimeout   int64

	// nerves are channels transferring information in/out brain
	hearingNerve  chan []byte
	speakingNerve chan []byte
}

// initialization
func (tcpTalk *TCPTalk) init() {
	// heartbeat
	tcpTalk.interval = DefaultHeartBeatInterval

	// timeout
	tcpTalk.connectTimeout = DefaultConnectTimeout
	tcpTalk.speakTimeout = DefaultSpeakTimeout

	// nerve
	tcpTalk.hearingNerve = make(chan []byte, DefaultNerveBuffer)
	tcpTalk.speakingNerve = make(chan []byte, DefaultNerveBuffer)
}

// Gossip implement of speak to
func (tcpTalk *TCPTalk) Gossip(member *Member, messages []byte) (echo []byte, err error) {
	return
}

// Hear implement of hear
func (tcpTalk *TCPTalk) Hear() (err error) {
	return
}

// Brain implement of brain
func (tcpTalk *TCPTalk) Brain() (err error) {
	return
}
