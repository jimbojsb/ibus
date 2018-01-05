package main

import (
	"encoding/hex"
	"github.com/jacobsa/go-serial/serial"
	"os"
	"strings"
)

func main() {
	ttyPath := os.Args[1]

	options := serial.OpenOptions{
		PortName:              ttyPath,
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

	hexBytes := strings.Split(os.Args[2], " ")
	for _, hexByte := range hexBytes {
		binaryByte, _ := hex.DecodeString(hexByte)
		port.Write(binaryByte)
	}
}
