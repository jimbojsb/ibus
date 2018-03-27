package ibus

const DEVICE_RADIO = 0x68
const DEVICE_BOARD_MONITOR_BUTTONS = 0xf0
const DEVICE_IKE = 0x80
const DEVICE_MFW = 0x50
const DEVICE_NAV_COMPUTER = 0x30
const DEVICE_PARK_DISTANCE = 0x60
const DEVICE_LIGHT_CONTROL = 0xbf
const DEVICE_NAV_LOCATION = 0xd0
const DEVICE_BROADCAST = 0xff
const DEVICE_CDPLAYER = 0x18

var (
	CdPlayer = &Device_CdPlayer{}
	Ike      = &Device_Ike{}
	Radio    = &Device_Radio{}
)
