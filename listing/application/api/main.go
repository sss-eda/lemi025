package main

import (
	"log"
	"net/http"
)

// This looks good. But now we need to think about the configuration...
//   --> Now this is where we start needing structs and just referring to the
//       packages becomes useless.
//   --> We need to have configurable structs if we want to be able to
//       dynamically change adapters.

func main() {
	http.HandleFunc(
		"lemi025/instrument/{id}", // GET method
		rest.GetInstrumentByIDHandler( // returns http.HandlerFunc
			listing.GetInstrumentByIDUseCase( // returns func(getinstrumentbyid.request) getinstrumentbyid.response
				postgres.GetInstrumentByIDRepo(db), // returns func(instrument.id) instrument
			),
		),
	)

	http.HandleFunc(
		"lemi025/instruments", // GET method
		rest.ListInstrumentsHandler( // returns http.HandlerFunc
			listing.ListInstrumentsUseCase( // returns func(request) response
				postgres.ListInstrumentsRepo(db), // returns func(id) instrument
			),
		),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
