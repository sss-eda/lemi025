package serial

import (
	"io"

	"github.com/sss-eda/lemi025/internal/domain"
)

// CheckFLASHData TODO
func CheckFLASHData(w io.Writer) func(domain.CheckFLASHDataCommand) error {
	return func(command domain.CheckFLASHDataCommand) error {
		_, err := w.Write([]byte{0x3D, 0x3A})
		if err != nil {
			return err
		}

		return nil
	}
}
