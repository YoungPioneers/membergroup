// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"sync"
)

// MemberStateType type of member state
type MemberStateType int

const (
	// DefaultPort .
	DefaultPort uint32 = 1314

	// MemberAlive indicates member is alive
	MemberAlive MemberStateType = iota
	// MemberSuspect indicates member is suspected alive
	MemberSuspect
	// MemberDead indicates member is dead
	MemberDead
)

// Member single member in cluster
type Member struct {
	uid int64

	// ipv4
	ip string
	// all members should use an uniform port for communications
	port uint32

	// member status
	status MemberStateType

	// talk
	talk Talk

	// nerves are channels transferring information in/out brain
	hearingNerve  chan []byte
	speakingNerve chan []byte

	lock *sync.RWMutex

	// heartbeat interval
	interval int64
}

// NewMember initialize a new member
func NewMember(talkType TalkType, ip string, port uint32) (member *Member, err error) {
	member = new(Member)
	member.uid = NewUID()
	member.ip = ip
	member.port = port
	member.status = MemberAlive

	// lock for read/write concurrently
	member.lock = new(sync.RWMutex)

	// nerve
	member.hearingNerve = make(chan []byte, DefaultNerveBuffer)
	member.speakingNerve = make(chan []byte, DefaultNerveBuffer)

	// init talk
	switch talkType {
	case TCPTalkType:
		member.talk = new(TCPTalk)
	case UDPTalkType:
		member.talk = new(UDPTalk)
	default:
		return nil, ErrTalkTypeNotDefined
	}
	err = member.talk.init(member, member.hearingNerve, member.speakingNerve)

	// heartbeat
	member.interval = DefaultHeartBeatInterval

	// begin to run
	go member.Brain()
	go member.Mouse()

	return member, err
}

// UID return unique id
func (member *Member) UID() int64 {
	member.lock.RLock()
	defer member.lock.RUnlock()
	return member.uid
}

// IP return ip address
func (member *Member) IP() string {
	member.lock.RLock()
	defer member.lock.RUnlock()
	return member.ip
}

// Port return the port
func (member *Member) Port() uint32 {
	member.lock.RLock()
	defer member.lock.RUnlock()
	return member.port
}

// Close cleaning up
func (member *Member) Close() (err error) {
	err = member.talk.Close()
	return
}

// Hearing .
func (member *Member) Hearing() bool {
	return member.Hearing()
}
