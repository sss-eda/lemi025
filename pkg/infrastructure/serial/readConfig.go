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

// OnConfigRead TODO
func OnConfigRead(
	strategy lemi025.OnConfigReadStrategy,
) func([5]byte) {
	return func(input [5]byte) {
		if string(input[:4]) != "L025 " {
			log.Printf("invalid input. Expexted L025, got: %s", input[:4])
		}

	}
}
