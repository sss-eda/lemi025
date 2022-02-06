package serial

import (
	"io"

	"github.com/sss-eda/lemi025/internal/domain"
)

// ReadCoefficients2 TODO
func ReadCoefficients2(
	w io.Writer,
) func(domain.ReadCoefficients2Command) error {
	return func(command domain.ReadCoefficients2Command) error {
		_, err := w.Write([]byte{0x3D, 0x36})
		if err != nil {
			return err
		}

		return nil
	}
}
