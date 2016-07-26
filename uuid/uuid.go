// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"net"
	"sync"
	"time"
)

type UUID [16]byte

const (
	dash       byte = '-'
	epochStart      = 122192928000000000
)

var (
	lastTime uint64
	lock     *sync.Mutex
)

func init() {
	lock = new(sync.Mutex)
}

func clockSequence() uint16 {
	buf := make([]byte, 2)
	safeRandom(buf)
	return binary.BigEndian.Uint16(buf)
}

func hardwareAddr() [6]byte {
	var hardwareAddr [6]byte
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			if len(iface.HardwareAddr) >= 6 {
				copy(hardwareAddr[:], iface.HardwareAddr)
				return hardwareAddr
			}
		}
	}

	safeRandom(hardwareAddr[:])

	hardwareAddr[0] |= 0x01
	return hardwareAddr
}

func safeRandom(dest []byte) {
	if _, err := rand.Read(dest); err != nil {
		panic(err)
	}
}

func machineInfo() (uint64, uint16, []byte) {
	lock.Lock()
	defer lock.Unlock()
	clockSequence := clockSequence()
	hardwareAddr := hardwareAddr()
	timeNow := epochStart + uint64(time.Now().UnixNano()/100)
	if timeNow <= lastTime {
		clockSequence++
	}
	lastTime = timeNow

	return timeNow, clockSequence, hardwareAddr[:]
}

func (u *UUID) SetVersion(v byte) {
	u[6] = (u[6] & 0x0f) | (v << 4)

}

// SetVariant sets variant bits as described in RFC 4122.
func (u *UUID) SetVariant() {
	u[8] = (u[8] & 0xbf) | 0x80
}

func (u UUID) String() string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u[0:4])
	buf[8] = dash
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = dash
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = dash
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = dash
	hex.Encode(buf[24:], u[10:])

	return string(buf)

}

// NewV1 returns UUID based on current timestamp and MAC address.
func NewV1() UUID {
	u := UUID{}

	timeNow, clockSeq, hardwareAddr := machineInfo()

	binary.BigEndian.PutUint32(u[0:], uint32(timeNow))
	binary.BigEndian.PutUint16(u[4:], uint16(timeNow>>32))
	binary.BigEndian.PutUint16(u[6:], uint16(timeNow>>48))
	binary.BigEndian.PutUint16(u[8:], clockSeq)

	copy(u[10:], hardwareAddr)

	u.SetVersion(1)
	u.SetVariant()

	return u
}
