// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"net"
)

// Member single member in cluster
type Member struct {
	Uid  int64
	Addr net.IP
	Port uint16
}

// NewMember initialize a new member
func NewMember() (member *Member) {
	member = new(Member)
	member.Uid = NewUid()

	// TODO init other attributes

	return member
}
