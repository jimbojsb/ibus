package ibus

type Device_Ike struct {
}

func (i *Device_Ike) WakeUp() {
	p := NewPacketWithoutChecksum(DEVICE_IKE, DEVICE_LIGHT_CONTROL, []byte{0x11, 0x02})
	WritePackets <- p
}
