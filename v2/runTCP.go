package main

import (
	"fmt"
	"log"
	"net"
)

func run() {
	const (
		host   = "localhost"
		port   = "8080"
		trType = "tcp"
	)

	listener, err := net.Listen(trType, fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Print(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	for {
		log.Println("Handling Request")
		buffer := make([]byte, 1024)

		length, err := conn.Read(buffer)
		if err != nil {
			log.Print(err)
		}

		str := string(buffer[:length])

		log.Println(conn.RemoteAddr().String())
		log.Printf("Received command %d\t:%q\n", length, str)

		send := func(res string, conn net.Conn) { conn.Write([]byte(res + "\n")) }
		switch str {
		case "PING\n":
			send("PONG", conn)
		case "PUSH\n":
			send("GOT PUSH", conn)
		case "QUIT\n":
			send("Goodbye", conn)
			conn.Close()
			return
		default:
			conn.Write([]byte(fmt.Sprintf("UNKNOWN_COMMAND: %s\n", str)))
		}
	}
}
