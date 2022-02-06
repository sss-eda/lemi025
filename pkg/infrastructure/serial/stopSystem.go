package serial

import (
	"io"

	"github.com/sss-eda/lemi025/internal/domain"
)

// StopSystem TODO
func StopSystem(w io.Writer) func(domain.StopSystemCommand) error {
	return func(command domain.StopSystemCommand) error {
		_, err := w.Write([]byte{0x3D, 0x38})
		if err != nil {
			return err
		}

		return nil
	}
}
