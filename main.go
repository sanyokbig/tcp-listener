package main

import (
	"github.com/sanyokbig/tcp-listener/ffxiv"
	"github.com/sanyokbig/tcp-listener/handler"
	"log"
)

const testPort = 15000

func main() {
	//mock.Run(testPort)

	//
	//go func() {
	//	log.Println(handler.ListenOnPort(testPort))
	//
	//}()

	go func() {
		log.Println(handler.ListenOnPort(ffxiv.GetPort()))
	}()

	<-make(chan struct{})
}
