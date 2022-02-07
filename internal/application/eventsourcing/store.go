package eventsourcing

// Store TODO
type Store interface {
	Load(ID) (Aggregate, error)
	Save(ID, Aggregate) error
}
