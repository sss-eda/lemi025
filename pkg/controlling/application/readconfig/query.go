package readconfig

import "log"

// Query TODO
func Query(
	readconfig func() error,
) func(func(*Response) error, *Request) {
	return func(w func(*Response) error, r *Request) {
		err := readconfig()
		err = w(&Response{
			Error: err,
		})
		if err != nil {
			log.Println(err)
		}
	}
}

// QueryFunc TODO
type QueryFunc[Response, Request any] func(func(*Response) error, *Request) error

// Request TODO
type Request struct{}

// Response TODO
type Response struct {
	Error error
}

// // Error TODO
// type Error struct {
// 	msg string
// }
