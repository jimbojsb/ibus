package ibus

import "bytes"

type Device_CdPlayer struct {
}

func (c *Device_CdPlayer) Handle(p *Packet) {
	switch p.Source {
	case DEVICE_RADIO:
		if bytes.Equal(p.Message, []byte{0x01}) {
			emit(EVENT_CDPLAYER_PING)
		} else if bytes.Equal(p.Message, []byte{0x38, 0x00, 0x00}) {
			emit(EVENT_CDPLAYER_STATUS)
		} else if bytes.Equal(p.Message, []byte{0x38, 0x03, 0x00}) {
			emit(EVENT_CDPLAYER_CONTROL_PLAY)
		} else if bytes.Equal(p.Message, []byte{0x38, 0x01, 0x00}) {
			emit(EVENT_CDPLAYER_CONTROL_STOP)
		} else if bytes.Equal(p.Message, []byte{0x38, 0x0a, 0x00}) {
			emit(EVENT_CDPLAYER_CONTROL_NEXT_TRACK)
		} else if bytes.Equal(p.Message, []byte{0x38, 0x0a, 0x01}) {
			emit(EVENT_CDPLAYER_CONTROL_PREVIOUS_TRACK)
		}
	}
}

func (c *Device_CdPlayer) Announce() {
	p := NewPacketWithoutChecksum(DEVICE_CDPLAYER, DEVICE_BROADCAST, []byte{0x02, 0x01})
	WritePackets <- p
}

func (c *Device_CdPlayer) Pong() {
	p := NewPacketWithoutChecksum(DEVICE_CDPLAYER, DEVICE_BROADCAST, []byte{0x02, 0x00})
	WritePackets <- p
}

func (c *Device_CdPlayer) RespondToStatusRequest(playing bool, disc int, track int) {
	var message []byte
	if playing {
		message = []byte{0x39, 0x00, 0x09, 0x00, 0x3f, 0x00}
	} else {
		message = []byte{0x39, 0x00, 0x02, 0x00, 0x3f, 0x00}
	}
	message = append(message, byte(disc))
	message = append(message, byte(track))
	p := NewPacketWithoutChecksum(DEVICE_CDPLAYER, DEVICE_RADIO, message)
	WritePackets <- p
}
