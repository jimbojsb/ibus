package ibus

const RADIO_MODEL_CD43 = 0
const RADIO_MODEL_CD53 = 1

type Device_Radio struct {
	Model int
}

func (r *Device_Radio) WriteText(text string) {
	p := NewPacketWithoutChecksum(0xc8, 0x80, []byte{0x23, 0x42, 0x01, 0x48, 0x65, 0x6c, 0x6c, 0x6f})
	WritePackets <- p
}
