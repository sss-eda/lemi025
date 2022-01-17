package eventsourcing

type EntityID interface {
	String() string
	Equals(EntityID) bool
}

type Entity interface {
	ID() EntityID
}
