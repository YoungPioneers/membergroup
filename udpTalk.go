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

	// nerves are channels transferring information in/out brain
	hearingNerve  chan []byte
	speakingNerve chan []byte

	// closed sign of closed
	closed bool
}

// initialization
func (udpTalk *UDPTalk) init() {
	// lock
	udpTalk.lock = new(sync.RWMutex)

	// heartbeat
	udpTalk.interval = DefaultHeartBeatInterval

	// nerve
	udpTalk.hearingNerve = make(chan []byte, DefaultNerveBuffer)
	udpTalk.speakingNerve = make(chan []byte, DefaultNerveBuffer)

	udpTalk.closed = false

	// begin to run
	go udpTalk.Brain()
	go udpTalk.Ear()
	go udpTalk.Mouse()
}

// Gossip implement of speak to
func (udpTalk *UDPTalk) Gossip(member *Member, messages []byte) (echo []byte, err error) {
	return
}

// Brain implement of brain
func (udpTalk *UDPTalk) Brain() (err error) {
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

// Mouse implement of mouse
func (udpTalk *UDPTalk) Mouse() (err error) {
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
