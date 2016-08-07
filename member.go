// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"sync"
)

// MemberStateType type of member state
type MemberStateType int

const (
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

	lock *sync.RWMutex
}

// NewMember initialize a new member
func NewMember() (member *Member) {
	member = new(Member)
	member.uid = NewUID()
	member.status = MemberAlive

	member.lock = new(sync.RWMutex)

	// TODO init other attributes

	return member
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
