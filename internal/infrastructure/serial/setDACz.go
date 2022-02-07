package serial

import (
	"encoding/binary"
	"io"

	"github.com/sss-eda/lemi025/internal/domain"
)

// SetDACz TODO
func SetDACz(w io.Writer) func(domain.SetDACzCommand) error {
	return func(command domain.SetDACzCommand) error {
		buffer := make([]byte, 4)

		buffer = []byte{0x3D, 0x3F}
		binary.LittleEndian.PutUint16(buffer[2:4], uint16(command.Value))

		_, err := w.Write(buffer)
		if err != nil {
			return err
		}

		return nil
	}
}
