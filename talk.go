// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

// Support propagation with tcp/udp protocol using an uniform prot, default 985.
// Each instance may contain propagate serveral client and one server.

const (
	// DefaultRTT 2 millisecond for rtt
	DefaultRTT = 2
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
	// SpeakTo say something to other
	SpeakTo(member *Member, messages []byte) (echo []byte, err error)
	// Hear listen from others
	Hear() (err error)
	// Brain deal with messages heard from others and control speaks to others
	Brain() (err error)
}
