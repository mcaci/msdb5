package main

import (
	"flag"

	"github.com/mcaci/msdb5/v2/frw"
)

func main() {
	noSide := flag.Bool("no-side", false, "Add flag to specify no side deck is to be used.")
	flag.Parse()

	frw.Game(*noSide)
	// frw.Run()
}
