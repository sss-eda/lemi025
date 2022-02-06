package serial

import (
	"io"

	"github.com/sss-eda/lemi025/internal/domain"
)

// ReadGPSData TODO
func ReadGPSData(w io.Writer) func(domain.ReadGPSDataCommand) error {
	return func(command domain.ReadGPSDataCommand) error {
		_, err := w.Write([]byte{0x3D, 0x37, 0xFF, 0xFF})
		if err != nil {
			return err
		}

		return nil
	}
}
