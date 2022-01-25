package lemi025

import "io"

// Service TODO
type Service struct {
	connection io.ReadWriteCloser
}

// ReadConfig TODO
func (service *Service) ReadConfig() error {
	return nil
}

// ReadTime TODO
func (service *Service) ReadTime() error {
	return nil
}

// SetTime TODO
func (service *Service) SetTime() error {
	return nil
}
