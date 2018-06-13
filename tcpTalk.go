// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	"bufio"
	"fmt"
	"net"
	"sync"

	log "github.com/YoungPioneers/blog4go"
)

// TCPTalk propagate with tcp protocol
type TCPTalk struct {
	// lock .
	lock *sync.RWMutex

	// timeout
	connectTimeout int64
	speakTimeout   int64

	// port
	hearingPort uint32

	// nerves
	hearingNerve  chan []byte
	speakingNerve chan []byte

	// status
	hearing bool

	// closed sign of closed
	closed bool
	// closeChn
	closeChn chan bool
}

// init initialization linking hearingNerve and speakingNerve
func (tcpTalk *TCPTalk) init(own *Member, hearingPort uint32, hearingNerve, speakingNerve chan []byte) (err error) {
	log.Debug("tcp talk init")
	// lock
	tcpTalk.lock = new(sync.RWMutex)

	// timeout
	tcpTalk.connectTimeout = DefaultConnectTimeout
	tcpTalk.speakTimeout = DefaultSpeakTimeout

	// port
	tcpTalk.hearingPort = hearingPort

	// nerve
	tcpTalk.hearingNerve = hearingNerve
	tcpTalk.speakingNerve = speakingNerve

	// status
	tcpTalk.hearing = false

	tcpTalk.closed = false
	tcpTalk.closeChn = make(chan bool)

	// begin to run
	go tcpTalk.Ear()

	for !tcpTalk.Hearing() {
	}

	return
}

// Gossip implement of speak to
func (tcpTalk *TCPTalk) Gossip(ip string, port uint32, message []byte) (echo []byte, err error) {
	log.Debugf("Gossip to %s:%d, message: %s\n", ip, port, message)

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, port))
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}

	defer conn.Close()
	log.Debug("connected!")

	conn.Write(message)
	conn.Write([]byte{TalkDelimeter})
	log.Debugf("write message: %s", message)

	reader := bufio.NewReader(conn)
	echo, err = reader.ReadBytes(TalkDelimeter)
	if nil != err {
		log.Error(err.Error())
		return nil, err
	}
	log.Debugf("echo: %s", echo)

	return echo, err
}

// Ear implement of ear
func (tcpTalk *TCPTalk) Ear() (err error) {
	// run a tcp server to receive messages
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", tcpTalk.hearingPort))
	if nil != err {
		log.Error(err.Error())
		return
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if nil != err {
		log.Error(err.Error())
		return
	}
	defer tcpListener.Close()

	log.Debug("ear is hearing")
	tcpTalk.lock.Lock()
	tcpTalk.hearing = true
	tcpTalk.lock.Unlock()
EarLoop:
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if nil != err {
			log.Error(err.Error())
			continue
		}
		log.Debugf("A client connected: %s", tcpConn.RemoteAddr().String())

		if tcpTalk.Closed() {
			break EarLoop
		}

		err = tcpTalk.Hear(tcpConn)
		if nil != err {
			log.Error(err.Error())
			continue
		}
	}

	return
}

// Hear implement of hear
func (tcpTalk *TCPTalk) Hear(tcpConn *net.TCPConn) (err error) {
	log.Debug("let me hear")
	reader := bufio.NewReader(tcpConn)
	bytes, err := reader.ReadBytes(TalkDelimeter)
	if nil != err {
		// response remote
		log.Errorf("hear fail. err: %s", err.Error())
		tcpConn.Write([]byte(TalkFailResponse))
		tcpConn.Write([]byte{TalkDelimeter})
		return
	}
	log.Debugf("hear success, %s", bytes)

	// response remote
	tcpConn.Write([]byte(TalkSuccessResponse))
	tcpConn.Write([]byte{TalkDelimeter})

	// notify my brain
	tcpTalk.hearingNerve <- bytes
	tcpConn.Close()

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
	tcpTalk.lock.RLock()
	defer tcpTalk.lock.RUnlock()

	return tcpTalk.closed
}

// Hearing .
func (tcpTalk *TCPTalk) Hearing() bool {
	tcpTalk.lock.RLock()
	defer tcpTalk.lock.RUnlock()

	return tcpTalk.hearing
}
