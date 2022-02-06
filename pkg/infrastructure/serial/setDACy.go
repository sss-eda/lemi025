package serial

import (
	"encoding/binary"
	"io"

	"github.com/sss-eda/lemi025/internal/domain"
)

// SetDACy TODO
func SetDACy(w io.Writer) func(domain.SetDACyCommand) error {
	return func(command domain.SetDACyCommand) error {
		buffer := make([]byte, 4)

		buffer = []byte{0x3D, 0x3E}
		binary.LittleEndian.PutUint16(buffer[2:4], uint16(command.Value))

		_, err := w.Write(buffer)
		if err != nil {
			return err
		}

		return nil
	}
}
