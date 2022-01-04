package commanding

import (
	"io"

	"github.com/sss-eda/lemi025"
)

// Service TODO
type Service struct {
	io.Writer
}

// ReadConfig TODO
func (service *Service) ReadConfig(
	command lemi025.ReadConfigCommand,
) error {
	_, err := service.Write([]byte{0x3D, 0x30})
	if err != nil {
		return err
	}

	return nil
}
