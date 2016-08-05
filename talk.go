// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

// Support propagation with tcp/udp protocol using an uniform prot, default 985.
// Each instance may contain propagate serveral client and one server.

const (
	// DefaultRTT 2 millisecond for rtt, suspected timeout maybe 3 * RTT
	DefaultRTT = 2
	// DefaultK choose k members in every gossip, depends on group size
	DefaultK = 1

	// DefaultConnectTimeout 10 millisecond for connection timeout
	DefaultConnectTimeout = 10
	// DefaultSpeakTimeout 100 millisecond waiting for speak response
	DefaultSpeakTimeout = 5
	// DefaultHeartBeatInterval send heart beat every 1000 millisecond
	DefaultHeartBeatInterval = 1000

	// DefaultNerveBuffer determines buffer size for nerve channels
	DefaultNerveBuffer = 1024
)

// Talk is base of propagation
type Talk interface {
	// TODO echo use an struct
	// Gossip spread information to others
	Gossip(member *Member, messages []byte) (echo []byte, err error)
	// Hear listen information from others
	Hear() (err error)
	// Brain deal with messages heard from others and control speaks to others
	Brain() (err error)
	// Detection failure detection
	Dection() (err error)
	// Ping .
	Ping(member *Member) (echo []byte, err error)
	// PingReq ask others to ping
	PingReq(friend *Member, target *Member) (echo []byte, err error)
}
