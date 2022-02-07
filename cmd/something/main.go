package main

import (
	"log"

	"github.com/sss-eda/lemi025"
)

// This is what we want our main func to look like eventually I think.
func main() {
	instrument, err := lemi025.New(
		lemi025.JetstreamStore,
		lemi025.SerialStrategy,
		lemi025.JetstreamAPI,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(instrument.Drive())
}
