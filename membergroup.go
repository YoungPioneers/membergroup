// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

// Package membergroup is an implement of SWIM alorithm
// It aims at maintaining status of members in a decentralized cluster and
// providing easy ways to communicate amoung all members
package membergroup

// MemberGroup groups of members in cluster
type MemberGroup struct {
	// key: uid
	members map[int64]*Member
}

// Create create a member group
func Create() (group *MemberGroup) {
	group = new(MemberGroup)
	group.members = make(map[int64]*Member)

	// 将自身加入到组中
	own := NewMember()
	group.members[own.UID()] = own

	return group
}

// Members return members
func (membergroup *MemberGroup) Members() map[int64]*Member {
	return membergroup.members
}

// Join join a member group
func (membergroup *MemberGroup) Join(err error) {
	return
}
