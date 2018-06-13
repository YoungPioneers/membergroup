// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

import (
	log "github.com/YoungPioneers/blog4go"
)

// Brain deal with messages heard from others and control speaks to others
func (member *Member) Brain() (err error) {
	log.Debug("start thinking")
	for {
		select {
		case bytes := <-member.hearingNerve:
			// brain think abount things from hearing nerve
			// then output to speaking nerve
			var heardMessage Message
			err := decode(bytes, heardMessage)
			if nil != err {
				log.Errorf("decode err: %s", err.Error())
				continue
			}

			log.Debugf("message: %v", heardMessage)
			// some logic work of brain
			gossipMessage := heardMessage
			gossipBytes, err := encode(gossipMessage)
			if nil != err {
				log.Errorf("encode err: %s", err.Error())
				continue
			}

			member.speakingNerve <- gossipBytes
		}
	}
}
