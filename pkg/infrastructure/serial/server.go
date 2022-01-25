package serial

import "io"

// Server TODO
type Server struct {
	service Service
}

// Serve TODO
func (server *Server) Serve(
	r io.Reader,
) error {

}
