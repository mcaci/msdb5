package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/mcaci/msdb5/v2/app/game"
	"github.com/mcaci/msdb5/v2/frw"
)

func main() {
	noSide := flag.Bool("no-side", false, "Flag to specify no side deck is to be used. Default: false (no side deck used)")
	host := flag.String("host", "localhost", "Hostname. Default: localhost.")
	port := flag.String("port", "8080", "Port number. Default: 8080.")
	network := flag.String("network", "tcp", "Network protocol. Default: tcp.")
	flag.Parse()

	// setup game
	_ = game.NewGame(&game.Options{
		WithSide: !*noSide,
	})

	// game.WaitForPlayers(g, listen.WithAINames)
	// game.Start(g)
	//
	// log.Println("Match over", g)
	// log.Println("Score", game.Score(g))

	// setup connection listenner
	l, err := net.Listen(*network, fmt.Sprintf("%s:%s", *host, *port))
	if err != nil {
		log.Print(err)
		return
	}
	defer l.Close()

	// accept connection from players
	for {
		c, err := l.Accept()
		if err != nil {
			log.Print(err)
		}
		go func() {
			_, e := frw.Handle(c)
			for e == nil {
				_, e = frw.Handle(c)
			}
		}()
	}
}
