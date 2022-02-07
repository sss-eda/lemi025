package serial

import (
	"io"

	"github.com/sss-eda/encoding/bcd"

	"github.com/sss-eda/lemi025/internal/domain"
)

// SetCoefficients1 TODO
func SetCoefficients1(
	w io.Writer,
) func(domain.SetCoefficients1Command) error {
	return func(command domain.SetCoefficients1Command) error {
		mode, err := bcd.Encode(byte(command.Mode))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte{0x3D, 0x33, 0xFF, mode})
		if err != nil {
			return err
		}

		return nil
	}
}
