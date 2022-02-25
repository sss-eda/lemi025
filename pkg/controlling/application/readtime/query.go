package readtime

import "log"

// Query TODO
func Query(
	readtime func() error,
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

// Request TODO
type Request struct{}

// Response TODO
type Response struct {
	Error error
}
