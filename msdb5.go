package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/mcaci/msdb5/catalog"
	"github.com/mcaci/msdb5/frw"
	"golang.org/x/text/language"
)

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	var noSide = flag.Bool("no-side", false, "Add flag to specify no side deck is to be used.")
	var lang = flag.String("lang", "it", "Add flag to specify the language to use.")
	flag.Parse() // parse the flags

	langTag := language.MustParse(*lang)
	r := frw.NewGameRoom(!*noSide, langTag)
	http.Handle("/", frw.NewTemplateHandler(langTag))
	http.Handle("/room", r)

	// get the room going
	go r.Run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
