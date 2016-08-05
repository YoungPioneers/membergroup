// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

// UDPTalk propagate with udp protocol
type UDPTalk struct {
	// heart beat interval
	interval int64

	// nerves are channels transferring information in/out brain
	hearingNerve  chan []byte
	speakingNerve chan []byte
}

// initialization
func (udpTalk *UDPTalk) init() {
	// heartbeat
	udpTalk.interval = DefaultHeartBeatInterval

	// nerve
	udpTalk.hearingNerve = make(chan []byte, DefaultNerveBuffer)
	udpTalk.speakingNerve = make(chan []byte, DefaultNerveBuffer)
}

// Gossip implement of speak to
func (udpTalk *UDPTalk) Gossip(member *Member, messages []byte) (echo []byte, err error) {
	return
}

// Hear implement of hear
func (udpTalk *UDPTalk) Hear() (err error) {
	return
}

// Brain implement of brain
func (udpTalk *UDPTalk) Brain() (err error) {
	return
}
