package serial

import (
	"io"

	"github.com/sss-eda/lemi025"
)

// ReadConfig TODO
func ReadConfig(w io.Writer) func(lemi025.ReadConfigInput) error {
	return func(input lemi025.ReadConfigInput) error {
		_, err := w.Write([]byte{0x3D, 0x30})
		if err != nil {
			return err
		}

		return nil
	}
}
