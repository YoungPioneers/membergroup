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
	// UIDSignBits 1 bit for sign
	UIDSignBits = 1
	// UIDTimeBits 41 bits for milliseconds since custom epoch
	UIDTimeBits = 41
	// UIDWorkerBits 4 bits for worker id
	UIDWorkerBits = 6
	// UIDSeqBits 10 bits for sequence number
	UIDSeqBits = 16
)

// randomize bytes
func random(bytes []byte) {
	rand.Read(bytes)
}

// timeStamp
func timeStamp() int64 {
	ts := time.Now().Unix()

	// reserve 41 bits
	ts = ts & (1<<UIDTimeBits - 1)
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
	workID = workID & (1<<UIDWorkerBits - 1)

	return workID
}

// sequence id
func sequence() int64 {
	bytes := make([]byte, UIDSeqBits)
	random(bytes)
	// reserve 10 bits
	return int64(binary.BigEndian.Uint64(bytes)) & (1<<UIDSeqBits - 1)
}

// NewUID get an uniq id
func NewUID() int64 {
	var uid int64
	uid = uid | timeStamp()
	uid = (uid << UIDWorkerBits) | workID()
	uid = (uid << UIDSeqBits) | sequence()

	return uid
}
