package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/nikiforosFreespirit/msdb5/ui"
)

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	var side = flag.Bool("side", false, "Whether to use side deck or not.")
	flag.Parse() // parse the flags

	r := ui.NewRoom(*side)
	http.Handle("/", ui.NewTemplateHandler())
	http.Handle("/room", r)

	// get the room going
	go r.Run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
