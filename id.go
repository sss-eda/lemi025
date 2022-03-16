package lemi025

// ID TODO
type ID interface {
	Equals(ID) bool
	String() string
}
