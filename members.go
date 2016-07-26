// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package members

import (
	"./uuid"
	"net"
	"os"
)

type Member struct {
	Name string
	Addr net.IP
	Port uint16
}

type MemberGroup struct {
	members []*Member
}

func Create() (group *MemberGroup) {
	group = new(MemberGroup)
	group.members = make([]*Member, 0)

	// 将自身加入到组中
	own := new(Member)
	own.Name = uuid.NewV1().String()

	return group
}
