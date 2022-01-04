package observing

import "io"

// Service TODO
type Service struct {
	io.Reader
}

func (service *Service) Observe(
	ctx context.Context,

) {

}
