package readtime

// Query TODO
type Query interface {
	Request() *Request
	Respond(*Response) error
}
