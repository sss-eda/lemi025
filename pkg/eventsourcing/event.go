package eventsourcing

// Event TODO
type Event interface {
	Type() string
	Apply(*Entity)
}
