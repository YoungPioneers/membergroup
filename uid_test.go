// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"testing"
)

func TestUid(t *testing.T) {
	uid := NewUid()

	// first bit must be 0
	if uid>>63 != 0 {
		t.Error("uid first bit is not 0")
	}
}
