package ibus

import (
	"github.com/asaskevich/EventBus"
	"github.com/jacobsa/go-serial/serial"
	"sync"
)

var (
	ReadPackets  = make(chan *Packet, 16)
	WritePackets = make(chan *Packet, 16)
	Events       = EventBus.New()
)

func Run(device string) {

	options := serial.OpenOptions{
		PortName:              device,
		BaudRate:              9600,
		DataBits:              8,
		StopBits:              1,
		ParityMode:            serial.PARITY_EVEN,
		RTSCTSFlowControl:     true,
		MinimumReadSize:       1,
		InterCharacterTimeout: 0,
	}
	port, err := serial.Open(options)

	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		readBuffer := make([]byte, 0)

		shiftBuffer := func() {
			readBuffer = readBuffer[1:]
		}

		for {
			buf := make([]byte, 1)
			port.Read(buf)
			readBuffer = append(readBuffer, buf[0])

			if len(readBuffer) >= 5 {
				possibleLength := int(readBuffer[1])
				if possibleLength < 3 { // too small for a valid packet
					shiftBuffer()
					continue
				} else if possibleLength > 64 { // too big for a valid packet
					shiftBuffer()
					continue
				} else if len(readBuffer) < (2 + possibleLength) { // might be valid, but we don't have all the bytes yet
					continue
				} else if BytesArePacket(readBuffer) {
					p := NewPacketFromBytes(readBuffer)
					ReadPackets <- p
					Events.Publish(EVENT_PACKET_RECEIVED, p)
					readBuffer = readBuffer[2+possibleLength:]
				} else {
					shiftBuffer() // something else is wrong, who knows what
				}
			}
		}
	}()

	go func() {
		for {
			p := <-WritePackets
			_, err := port.Write(p.asByteSlice())
			if err != nil {
				panic(err)
			}
			Events.Publish(EVENT_PACKET_SENT, p)
		}
	}()

	go func() {
		for {
			p := <-ReadPackets
			route(p)
		}
	}()

	wg.Wait()
}
