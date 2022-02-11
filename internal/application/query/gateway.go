package query

// Gateway TODO
type Gateway func(Request) (Response, error)

// Handler TODO
type Handler func(Gateway) error
