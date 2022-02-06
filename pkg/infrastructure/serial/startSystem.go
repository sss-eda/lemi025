package serial

import (
	"io"

	"github.com/sss-eda/lemi025/internal/domain"
)

// StartSystem TODO
func StartSystem(w io.Writer) func(domain.StartSystemCommand) error {
	return func(command domain.StartSystemCommand) error {
		_, err := w.Write([]byte{0x3D, 0x39})
		if err != nil {
			return err
		}

		return nil
	}
}
