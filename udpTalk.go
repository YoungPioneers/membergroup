// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"sync"
)

// UDPTalk propagate with udp protocol
type UDPTalk struct {
	// lock .
	lock *sync.RWMutex

	// heart beat interval
	interval int64

	// port
	hearingPort uint32

	// nerves are channels transferring information in/out brain
	hearingNerve  chan []byte
	speakingNerve chan []byte

	// closed sign of closed
	closed bool
}

// initialization
func (udpTalk *UDPTalk) init(own *Member, hearingPort uint32, hearingNerve, speakingNerve chan []byte) (err error) {
	// lock
	udpTalk.lock = new(sync.RWMutex)

	// heartbeat
	udpTalk.interval = DefaultHeartBeatInterval

	// port
	udpTalk.hearingPort = hearingPort

	// nerve
	udpTalk.hearingNerve = make(chan []byte, DefaultNerveBuffer)
	udpTalk.speakingNerve = make(chan []byte, DefaultNerveBuffer)

	udpTalk.closed = false
	go udpTalk.Ear()

	return
}

// Gossip implement of speak to
func (udpTalk *UDPTalk) Gossip(ip string, port uint32, messages []byte) (echo []byte, err error) {
	return
}

// Ear implement of ear
func (udpTalk *UDPTalk) Ear() (err error) {
	return
}

// Hear implement of hear
func (udpTalk *UDPTalk) Hear() (err error) {
	return
}

// Ping send ping to other
func (udpTalk *UDPTalk) Ping(member *Member) (echo []byte, err error) {
	return
}

// PingReq send ping req to other via your friend
func (udpTalk *UDPTalk) PingReq(friend *Member, target *Member) (echo []byte, err error) {
	return
}

// Close close
func (udpTalk *UDPTalk) Close() (err error) {
	if udpTalk.closed {
		return ErrTalkAlreadyClosed
	}

	udpTalk.lock.Lock()
	defer udpTalk.lock.Unlock()

	udpTalk.closed = true
	return
}
