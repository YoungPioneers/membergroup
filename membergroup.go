// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

// Package membergroup is an implement of SWIM alorithm
// It aims at maintaining status of members in a decentralized cluster and
// providing easy ways to communicate amoung all members
package membergroup

import (
	"fmt"
	"sync"
)

// MemberGroup groups of members in cluster
type MemberGroup struct {
	own *Member
	// key: uid
	members map[int64]*Member

	lock *sync.RWMutex
}

// Create create a member group
func Create(talkType TalkType) (group *MemberGroup, err error) {
	group = new(MemberGroup)
	group.members = make(map[int64]*Member)

	// 将自身加入到组中
	ip := getLocalIP()
	group.own, err = NewMember(talkType, ip, DefaultPort)
	if nil != err {
		return nil, err
	}
	group.members[group.own.UID()] = group.own

	group.lock = new(sync.RWMutex)

	return group, err
}

// Members return members
func (membergroup *MemberGroup) Members() map[int64]*Member {
	membergroup.lock.RLock()
	defer membergroup.lock.RUnlock()
	return membergroup.members
}

// setMember set members
func (membergroup *MemberGroup) setMembers(members map[int64]*Member) {
	membergroup.lock.Lock()
	defer membergroup.lock.Unlock()
	membergroup.members = members
}

// Join join a member group
func (membergroup *MemberGroup) Join(addrs []string) error {
	for _, addr := range addrs {
		fmt.Printf("Join %s\n", addr)
		membergroup.own.talk.Gossip(addr, DefaultPort, []byte("hello"))
	}
	return nil
}

// Broadcast broadcase message among its member group
func (membergroup *MemberGroup) Broadcast(message []byte) error {
	return nil
}
