package instrument

// ID TODO
type ID interface {
	Equals(ID) bool
	String() string
}
