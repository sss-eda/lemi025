package id

// Identifiable TODO
type Identifiable interface {
	Equals(Identifiable) bool
}
