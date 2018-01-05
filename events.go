package ibus

const EVENT = "event"

const EVENT_PACKET_RECEIVED = "packet:received"
const EVENT_PACKET_SENT = "packet:sent"

const EVENT_CDPLAYER_PING = "cdplayer:ping"
const EVENT_CDPLAYER_STATUS_REQUEST = "cdplayer:status_request"

func emit(event string, args ...interface{}) {
	Events.Publish(EVENT, event)
	Events.Publish(event, args)
}
