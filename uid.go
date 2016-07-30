// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"crypto/rand"
	"encoding/binary"
	"net"
	"time"
)

// We generate unique id for each instance base on twitter's snow flake

const (
	// UidSignBits 1 bit for sign
	UidSignBits = 1
	// UidTimeBits 41 bits for milliseconds since custom epoch
	UidTimeBits = 41
	// UidWorkerBits 4 bits for worker id
	UidWorkerBits = 6
	// UidSeqBits 10 bits for sequence number
	UidSeqBits = 16
)

// randomize bytes
func random(bytes []byte) {
	rand.Read(bytes)
}

// timeStamp
func timeStamp() int64 {
	ts := time.Now().Unix()

	// reserve 41 bits
	ts = ts & (1<<UidTimeBits - 1)
	return ts
}

// workID
// use first mac as work id and 10 bits reserved
func workID() int64 {
	var workID int64
	bytes := make([]byte, 8)

	if interfaces, err := net.Interfaces(); nil == err {
		for _, iface := range interfaces {
			if length := len(iface.HardwareAddr); length >= 6 {
				copy(bytes, iface.HardwareAddr[:])
				break
			}
		}
	} else {
		// if get mac address failed, use random bytes
		random(bytes)
	}

	workID = int64(binary.BigEndian.Uint64(bytes))
	// reserve 10 bits
	workID = workID & (1<<UidWorkerBits - 1)

	return workID
}

// sequence id
func sequence() int64 {
	bytes := make([]byte, UidSeqBits)
	random(bytes)
	// reserve 10 bits
	return int64(binary.BigEndian.Uint64(bytes)) & (1<<UidSeqBits - 1)
}

// NewUid get an uniq id
func NewUid() int64 {
	var uid int64
	uid = uid | timeStamp()
	uid = (uid << UidWorkerBits) | workID()
	uid = (uid << UidSeqBits) | sequence()

	return uid
}
