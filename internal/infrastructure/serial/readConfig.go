package serial

import (
	"io"

	"github.com/sss-eda/lemi025/internal/domain"
)

// ReadConfig TODO
func ReadConfig(w io.Writer) func(domain.ReadConfigCommand) error {
	return func(command domain.ReadConfigCommand) error {
		_, err := w.Write([]byte{0x3D, 0x30, 0xFF, 0xFF})
		if err != nil {
			return err
		}

		return nil
	}
}
