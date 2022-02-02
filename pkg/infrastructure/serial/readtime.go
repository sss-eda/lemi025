package serial

import (
	"io"

	"github.com/sss-eda/lemi025"
)

// ReadTime TODO
func ReadTime(w io.Writer) func(lemi025.ReadTimeInput) error {
	return func(input lemi025.ReadTimeInput) error {
		_, err := w.Write([]byte{0x3D, 0x31})
		if err != nil {
			return err
		}

		return nil
	}
}
