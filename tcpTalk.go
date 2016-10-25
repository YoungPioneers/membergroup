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

	// timeout
	connectTimeout int64
	speakTimeout   int64

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
func (tcpTalk *TCPTalk) init(own *Member, hearingNerve, speakingNerve chan []byte) (err error) {
	fmt.Println("tcp talk init")
	// lock
	tcpTalk.lock = new(sync.RWMutex)

	// timeout
	tcpTalk.connectTimeout = DefaultConnectTimeout
	tcpTalk.speakTimeout = DefaultSpeakTimeout

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
	fmt.Printf("Gossip to %s, message: %s\n", fmt.Sprintf("%s:%d", ip, port), message)
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, port))
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}

	defer conn.Close()
	fmt.Println("connected!")

	conn.Write(message)
	conn.Write([]byte{TalkDelimeter})
	fmt.Printf("write message: %s\n", message)

	reader := bufio.NewReader(conn)
	echo, err = reader.ReadBytes(TalkDelimeter)
	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Printf("%s\n", echo)

	return echo, err
}

// Ear implement of ear
func (tcpTalk *TCPTalk) Ear() (err error) {

	// run a tcp server to receive messages
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", DefaultPort))
	if nil != err {
		fmt.Println(err.Error())
		return
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	defer tcpListener.Close()

	fmt.Println("ear is hearing")
	tcpTalk.hearing = true
EarLoop:
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if nil != err {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())

		if tcpTalk.Closed() {
			break EarLoop
		}

		err = tcpTalk.Hear(tcpConn)
		if nil != err {
			fmt.Println(err.Error())
			continue
		}
	}

	return
}

// Hear implement of hear
func (tcpTalk *TCPTalk) Hear(tcpConn *net.TCPConn) (err error) {
	fmt.Println("let me hear")
	reader := bufio.NewReader(tcpConn)
	bytes, err := reader.ReadBytes(TalkDelimeter)
	if nil != err {
		// response remote
		fmt.Println("hear fail")
		tcpConn.Write([]byte(TalkFailResponse))
		tcpConn.Write([]byte{TalkDelimeter})
		return
	}
	fmt.Printf("hear success, %s", bytes)

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
