package serial

import (
	"io"

	"github.com/sss-eda/lemi025/internal/domain"
)

// ReadCoefficients1 TODO
func ReadCoefficients1(
	w io.Writer,
) func(domain.ReadCoefficients1Command) error {
	return func(command domain.ReadCoefficients1Command) error {
		_, err := w.Write([]byte{0x3D, 0x34})
		if err != nil {
			return err
		}

		return nil
	}
}
