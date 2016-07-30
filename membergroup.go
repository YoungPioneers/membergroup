// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

// Package membergroup is an implement of SWIM alorithm
// It aims at maintaining status of members in a decentralized cluster and
// providing easy ways to communicate amoung all members
package membergroup

// MemberGroup groups of members in cluster
type MemberGroup struct {
	members []*Member
}

// Create create a member group
func Create() (group *MemberGroup) {
	group = new(MemberGroup)
	group.members = make([]*Member, 0)

	// 将自身加入到组中
	own := NewMember()

	group.members = append(group.members, own)

	return group
}
