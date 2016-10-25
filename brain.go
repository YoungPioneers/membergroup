// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

// Brain deal with messages heard from others and control speaks to others
func (member *Member) Brain() (err error) {
	for {
		select {
		case bytes := <-member.hearingNerve:
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

			member.speakingNerve <- gossipBytes
		}
	}
}
