// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

// Support propagation with tcp/udp protocol using an uniform prot, default 985.
// Each instance may contain propagate serveral client and one server.

import (
	"errors"
)

type TalkType int

const (
	// TCPTalkType tcp
	TCPTalkType TalkType = iota
	// UDPTalkType udp
	UDPTalkType

	// DefaultRTT 2 millisecond for rtt, suspected timeout maybe 3 * RTT
	DefaultRTT = 2
	// DefaultK choose k members in every gossip, depends on group size
	DefaultK = 1

	// DefaultConnectTimeout 10 millisecond for connection timeout
	DefaultConnectTimeout = 10
	// DefaultSpeakTimeout 100 millisecond waiting for speak response
	DefaultSpeakTimeout = 5
	// DefaultHeartBeatInterval send heart beat every 200 millisecond
	DefaultHeartBeatInterval = 200

	// DefaultNerveBuffer determines buffer size for nerve channels
	DefaultNerveBuffer = 1024

	// TalkDelimeter delimeter for talking protocol
	TalkDelimeter byte = '\n'
	// TalkSuccessResponse response for success
	TalkSuccessResponse string = "1"
	// TalkFailResponse response for failure
	TalkFailResponse string = "0"
)

var (
	ErrTalkTypeNotDefined = errors.New("Talk type not defined")
	ErrTalkAlreadyClosed  = errors.New("Talk had already closed")
)

// Talk is base of propagation
type Talk interface {
	// TODO echo use an struct
	// init  initialization linking hearingNerve and speakingNerve
	init(own *Member, hearingNerve, speakingNerve chan []byte) (err error)
	// Gossip spread information to others
	Gossip(ip string, port uint32, messages []byte) (echo []byte, err error)
	// Ear listen information from others
	Ear() (err error)
	// Hear receives information from ear
	//Hear(*net.TCPConn) (err error)
	// Detection failure detection
	//Dection() (err error)
	// Ping .
	Ping(member *Member) (echo []byte, err error)
	// PingReq ask others to ping
	PingReq(friend *Member, target *Member) (echo []byte, err error)
	// Close close
	Close() (err error)
}
