package rest

// ReadConfig TODO
func ReadConfig(
	mux.Router,
) {
	commandRouter.HandleFunc(
		"/commands/read-config",
		rest.HandleCommand( // returns http.HandlerFunc
			controlling.ExecuteCommand( // This might be usefull in terms of directing
				serial.ReadConfig(port), // returns
			),
		),
	).Methods("POST")
}
