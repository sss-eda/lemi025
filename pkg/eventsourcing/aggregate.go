package eventsourcing

type AggregateID interface {
	String() string
	Equals(AggregateID) bool
}

type AggregateType interface {
	String() string
	Equals(AggregateType) bool
}

// Aggregate TODO
type Aggregate interface {
	ID() AggregateID
	Type() AggregateType
	Root() Entity
	CommandHandler() func(Command) error
}
