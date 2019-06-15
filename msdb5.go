package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/nikiforosFreespirit/msdb5/frw"
)

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	var noSide = flag.Bool("no-side", false, "Add flag to specify no side deck is to be used.")
	flag.Parse() // parse the flags

	r := frw.NewGameRoom(!*noSide)
	http.Handle("/", frw.NewTemplateHandler())
	http.Handle("/room", r)

	// get the room going
	go r.Run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
