package readconfig

// Response TODO
type Response struct {
	Error error
}

// ResponseWriter TODO
type ResponseWriter func(*Response) error
