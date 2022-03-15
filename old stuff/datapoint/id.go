package datapoint

// ID TODO
type ID interface {
	Equals(ID) bool
	String() string
}
