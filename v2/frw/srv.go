package frw

import (
	"log"
	"net"
)

func Run(network, address string) {
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Print(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		go func() {
			for handle(conn) == nil {
			}
		}()
	}
}
