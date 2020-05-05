package mock

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func Run(port int) {
	go startListener(port)

	time.Sleep(time.Millisecond)

	go startSender(port)
}

func startListener(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}

	defer func() { _ = lis.Close() }()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("failed to accept conn: %v", conn)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		_, err := reader.ReadBytes('\n')
		if err != nil {
			log.Printf("failed to read bytes: %v", err)
			continue
		}
	}
}

func startSender(port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Printf("failed to open conn: %v", err)
	}

	writer := bufio.NewWriter(conn)

	for {
		_, err := writer.Write([]byte("Hello, World!\n"))
		if err != nil {
			log.Printf("failed to write: %v", err)
		}
		writer.Flush()

		time.Sleep(time.Second)

	}
}
