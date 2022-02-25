package datum

// Datum TODO
type Datum struct {
	x int
}

// Repository TODO
type Repository interface {
	Load(ID) (*Datum, error)
	Save(ID, Event) error
}

// Event TODO
type Event interface{}
