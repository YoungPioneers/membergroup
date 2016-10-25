// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

// Mouse implement of mouse
func (member *Member) Mouse() {
	//MouseLoop:
	for {
		select {
		case bytes := <-member.speakingNerve:
			// mouse gossip to other
			member := &Member{}
			echo, err := member.talk.Gossip("", 0, bytes)
			if nil != err {
				continue
			}
			member.hearingNerve <- echo
			//case <-member.closeChn:
			//break MouseLoop
		}
	}
}
