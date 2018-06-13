// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMemberGroup(t *testing.T) {
	Convey("Test tcp member group basic operations", t, func() {
		group, err := Create(TCPTalkType)
		So(err, ShouldBeNil)
		So(group, ShouldNotBeNil)
	})
}
