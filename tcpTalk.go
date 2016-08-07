// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

// TCPTalk propagate with tcp protocol
type TCPTalk struct {
	// lock .
	lock *sync.RWMutex

	// members
	own     *Member
	members *MemberGroup

	// heartbeat interval
	interval int64

	// timeout
	connectTimeout int64
	speakTimeout   int64

	// nerves are channels transferring information in/out brain
	hearingNerve  chan []byte
	speakingNerve chan []byte

	// closed sign of closed
	closed bool
	// closeChn
	closeChn chan bool
}

// initialization
func (tcpTalk *TCPTalk) init(own *Member) {
	// lock
	tcpTalk.lock = new(sync.RWMutex)

	// heartbeat
	tcpTalk.interval = DefaultHeartBeatInterval

	// timeout
	tcpTalk.connectTimeout = DefaultConnectTimeout
	tcpTalk.speakTimeout = DefaultSpeakTimeout

	// nerve
	tcpTalk.hearingNerve = make(chan []byte, DefaultNerveBuffer)
	tcpTalk.speakingNerve = make(chan []byte, DefaultNerveBuffer)

	tcpTalk.own = own

	tcpTalk.closed = false
	tcpTalk.closeChn = make(chan bool)

	// begin to run
	go tcpTalk.Brain()
	go tcpTalk.Ear()
	go tcpTalk.Mouse()
}

// Gossip implement of speak to
func (tcpTalk *TCPTalk) Gossip(member *Member, messages []byte) (echo []byte, err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", member.IP(), member.Port()))
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	reader := bufio.NewReader(conn)
	echo, err = reader.ReadBytes(TalkDelimeter)
	if nil != err {
		return
	}

	return
}

// Brain implement of brain
func (tcpTalk *TCPTalk) Brain() (err error) {
	for {
		select {
		case bytes := <-tcpTalk.hearingNerve:
			// brain think abount things from hearing nerve
			// then output to speaking nerve
			var heardMessage Message
			err := decode(bytes, heardMessage)
			if nil != err {
				continue
			}

			// some logic work of brain
			gossipMessage := heardMessage
			gossipBytes, err := encode(gossipMessage)
			if nil != err {
				continue
			}

			tcpTalk.speakingNerve <- gossipBytes
		}
	}
}

// Ear implement of ear
func (tcpTalk *TCPTalk) Ear() (err error) {
	// run a tcp server to receive messages
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", tcpTalk.own.IP(), tcpTalk.own.Port()))
	if nil != err {
		return
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if nil != err {
		return
	}
	defer tcpListener.Close()

EarLoop:
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if nil != err {
			continue
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())

		if tcpTalk.Closed() {
			break EarLoop
		}

		err = tcpTalk.Hear(tcpConn)
		if nil != err {
			continue
		}
	}

	return
}

// Hear implement of hear
func (tcpTalk *TCPTalk) Hear(tcpConn *net.TCPConn) (err error) {
	reader := bufio.NewReader(tcpConn)
	bytes, err := reader.ReadBytes(TalkDelimeter)
	if nil != err {
		// response remote
		tcpConn.Write([]byte(TalkFailResponse))
		return
	}

	// response remote
	tcpConn.Write([]byte(TalkSuccessResponse))

	// notify my brain
	tcpTalk.hearingNerve <- bytes
	tcpConn.Close()

	return
}

// Mouse implement of mouse
func (tcpTalk *TCPTalk) Mouse() (err error) {
MouseLoop:
	for {
		select {
		case bytes := <-tcpTalk.speakingNerve:
			// mouse gossip to other
			member := &Member{}
			echo, err := tcpTalk.Gossip(member, bytes)
			if nil != err {
				continue
			}
			tcpTalk.hearingNerve <- echo
		case <-tcpTalk.closeChn:
			break MouseLoop
		}
	}
	return
}

// Ping send ping to other
func (tcpTalk *TCPTalk) Ping(member *Member) (echo []byte, err error) {
	return
}

// PingReq send ping req to other via your frient
func (tcpTalk *TCPTalk) PingReq(friend *Member, target *Member) (echo []byte, err error) {
	return
}

// Close close
func (tcpTalk *TCPTalk) Close() (err error) {
	if tcpTalk.closed {
		return ErrTalkAlreadyClosed
	}

	tcpTalk.lock.Lock()
	defer tcpTalk.lock.Unlock()

	tcpTalk.closed = true
	return
}

// Closed closed
func (tcpTalk *TCPTalk) Closed() bool {
	return tcpTalk.closed
}
