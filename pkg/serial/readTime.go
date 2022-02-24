package serial

import (
	"io"
)

// ReadTime TODO
func ReadTime(writer io.Writer) error {
	_, err := writer.Write([]byte{0x3D, 0x31})
	return err
}

// TimeReaderFunc TODO
func TimeReaderFunc(
	writer io.Writer,
) func() error {
	return func() error {
		return ReadTime(writer)
	}
}
