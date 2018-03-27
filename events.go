package ibus

const EVENT = "event"

const EVENT_PACKET_RECEIVED = "packet:received"
const EVENT_PACKET_SENT = "packet:sent"

const EVENT_CDPLAYER_PING = "cdplayer:ping"
const EVENT_CDPLAYER_STATUS = "cdplayer:status"
const EVENT_CDPLAYER_CONTROL_PLAY = "cdplayer:control:play"
const EVENT_CDPLAYER_CONTROL_STOP = "cdstoper:control:stop"
const EVENT_CDPLAYER_CONTROL_NEXT_TRACK = "cdplayer:control:next_track"
const EVENT_CDPLAYER_CONTROL_PREVIOUS_TRACK = "cdplayer:control:previous_track"

func emit(event string, args ...interface{}) {
	Events.Publish(EVENT, event)
	if len(args) > 0 {
		Events.Publish(event, args)
	} else {
		Events.Publish(event)
	}

}
