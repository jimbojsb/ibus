package ibus

func route(p *Packet) {
	switch p.Destination {
	case IBUS_DEVICE_CDPLAYER:
		CdPlayer.Handle(p)
	}
}
