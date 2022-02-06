package serial

import (
	"io"

	"github.com/sss-eda/lemi025/internal/domain"
)

// ReadTime TODO
func ReadTime(w io.Writer) func(domain.ReadTimeCommand) error {
	return func(command domain.ReadTimeCommand) error {
		_, err := w.Write([]byte{0x3D, 0x31})
		if err != nil {
			return err
		}

		return nil
	}
}
