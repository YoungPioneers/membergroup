// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"github.com/golang/protobuf/proto"
)

func encode(message Message) (bytes []byte, err error) {
	return proto.Marshal(&message)
}

func decode(bytes []byte, message Message) (err error) {
	return proto.Unmarshal(bytes, &message)
}
