package frw

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
)

var errClientQuit = errors.New("received QUIT from client, ending connection")

func handle(c net.Conn) error {
	log.Println("Handling new Request")
	// read
	buffer := make([]byte, 1024)

	length, err := c.Read(buffer)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Println(c.RemoteAddr().String())
	log.Printf("Received command %d\t:%q\n", length, buffer[:length])
	s := string(bytes.TrimRight(buffer[:length], "\n"))

	// write
	write := func(s string) { c.Write([]byte(s + "\n")) }
	switch s {
	case "PING":
		write("PONG")
	case "PUSH":
		write("GOT PUSH")
	case "QUIT":
		write("Goodbye")
		c.Close()
		return errClientQuit
	default:
		write(fmt.Sprintf("UNKNOWN_COMMAND: %s\n", s))
	}
	return nil
}
