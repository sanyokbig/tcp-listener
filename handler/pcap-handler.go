package handler

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)


func ListenOnPort(port int) error {
	handler, err := pcap.OpenLive("any", 1600, true, 1000)
	if err != nil {
		panic(err)
	}

	defer handler.Close()

	err = handler.SetBPFFilter(fmt.Sprintf("tcp and port %v", port))
	if err != nil {
		panic(err)
	}

	source := gopacket.NewPacketSource(handler, handler.LinkType())

	for packet := range source.Packets() {
		handle(packet)
	}

	return nil
}

func handle(packet gopacket.Packet) {
	layer := packet.ApplicationLayer()
	if packet.ApplicationLayer() == nil {
		return
	}

	payload := layer.Payload()

	log.Printf("%v: %v", len(payload), payload)
}