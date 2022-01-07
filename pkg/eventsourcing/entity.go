package eventsourcing

// Entity TODO
type Entity struct {
	base   interface{}
	events []Event
}

// NewEntity TODO
func NewEntity(
	base interface{},
) *Entity {
	return &Entity{
		base:   base,
		events: []Event{},
	}
}

// On TODO
func (entity *Entity) On(event Event) {
	event.Apply(entity)
}
