package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/text/language"

	"github.com/nikiforosFreespirit/msdb5/frw"
)

type contextKey int

const (
	messagePrinterKey contextKey = 1
)

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	var noSide = flag.Bool("no-side", false, "Add flag to specify no side deck is to be used.")
	var lang = flag.String("lang", "it", "Add flag to specify the language to use.")
	flag.Parse() // parse the flags

	tag := language.Italian
	if *lang != "it" {
		tag = language.English
	}

	r := frw.NewGameRoom(!*noSide, tag)
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
