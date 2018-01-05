package main

import (
	"fmt"
	"github.com/jimbojsb/ibus"
	"os"
	"sync"
)

func main() {
	ttyPath := os.Args[1]
	fmt.Println("==Ibus Monitor==")

	ibus.Events.Subscribe(ibus.EVENT_PACKET_RECEIVED, func(p *ibus.Packet) {
		fmt.Println("<== " + p.String())
	})
	ibus.Events.Subscribe(ibus.EVENT_PACKET_SENT, func(p *ibus.Packet) {
		fmt.Println("==> " + p.String())
	})
	ibus.Events.Subscribe(ibus.EVENT, func(name string) {
		fmt.Println("<=> " + name)
	})

	wg := sync.WaitGroup{}
	wg.Add(1)
	go ibus.Run(ttyPath)
	wg.Wait()
}
