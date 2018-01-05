package ibus

import "encoding/hex"

type Packet struct {
	Source      byte
	Length      byte
	Destination byte
	Message     []byte
	Checksum    byte
}

func BytesArePacket(bytes []byte) bool {
	possibleChecksum := bytes[len(bytes)-1]
	calculatedChecksum := calculatePacketChecksum(bytes[0], bytes[1], bytes[2], bytes[3:len(bytes)-1])
	return possibleChecksum == calculatedChecksum
}

func NewPacketFromBytes(bytes []byte) *Packet {
	p := &Packet{}
	p.Source = bytes[0]
	p.Length = bytes[1]
	p.Destination = bytes[2]
	p.Message = bytes[3:(len(bytes) - 1)]
	p.Checksum = bytes[len(bytes)-1]
	return p
}

func NewPacketWithoutChecksum(source byte, destination byte, message []byte) *Packet {
	p := Packet{
		Source:      source,
		Length:      byte(len(message) + 2),
		Destination: destination,
		Message:     message,
	}
	p.setChecksum()
	return &p
}

func (p *Packet) asByteSlice() []byte {
	bytes := []byte{p.Source, p.Length, p.Destination}
	bytes = append(bytes, p.Message...)
	bytes = append(bytes, p.Checksum)
	return bytes
}

func (p *Packet) String() string {
	str := hex.EncodeToString([]byte{p.Source})
	str = str + " " + hex.EncodeToString([]byte{p.Length})
	str = str + " " + hex.EncodeToString([]byte{p.Destination})
	for _, el := range p.Message {
		str = str + " " + hex.EncodeToString([]byte{el})
	}
	str = str + " " + hex.EncodeToString([]byte{p.Checksum})
	return str
}

func (p *Packet) setChecksum() {
	p.Checksum = calculatePacketChecksum(p.Source, p.Length, p.Destination, p.Message)
}

func calculatePacketChecksum(source byte, length byte, destination byte, message []byte) byte {
	var xor byte
	xor = xor ^ source
	xor = xor ^ length
	xor = xor ^ destination
	for _, el := range message {
		xor = xor ^ el
	}
	return xor
}
