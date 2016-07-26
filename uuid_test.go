// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package members

import (
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := NewV1()
	t.Log(uuid.String())
}
