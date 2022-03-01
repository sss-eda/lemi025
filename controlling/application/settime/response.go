package settime

// Response TODO
type Response struct {
	Error error
}

// ResponseWriter TODO
type ResponseWriter func(*Response) error

// ResponseWriter TODO
// type ResponseWriter interface {
// 	Write(Response)
// 	Error(error)
// }
