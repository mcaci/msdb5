package frw

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

var errClientQuitMsdb5 = errors.New("Received QUIT from client, ending connection")

func Handle(c net.Conn) (struct{}, error) {
	log.Println("Handling new Request")
	// read
	buffer := make([]byte, 1024)

	length, err := c.Read(buffer)
	if err != nil {
		log.Print(err)
		return struct{}{}, err
	}

	log.Println(c.RemoteAddr().String())
	log.Printf("Received command %d\t:%q\n", length, buffer[:length])
	s := string(bytes.TrimRight(buffer[:length], "\n"))
	switch {
	case strings.HasPrefix(s, "join "):
		strings.Fields(s)
	}

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
		return struct{}{}, errClientQuitMsdb5
	default:
		write(fmt.Sprintf("UNKNOWN_COMMAND: %s\n", s))
	}
	return struct{}{}, nil
}
