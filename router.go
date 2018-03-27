package ibus

func route(p *Packet) {
	switch p.Destination {
	case DEVICE_CDPLAYER:
		CdPlayer.Handle(p)
	}
}
