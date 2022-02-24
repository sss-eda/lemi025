package serial

import (
	"io"
)

// ReadConfig TODO
func ReadConfig(writer io.Writer) error {
	_, err := writer.Write([]byte{0x3D, 0x30})
	return err
}

// ConfigReaderFunc TODO
func ConfigReaderFunc(
	writer io.Writer,
) func() error {
	return func() error {
		return ReadConfig(writer)
	}
}
