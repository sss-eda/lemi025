package serial

import "io"

func ReadConfig(
	w io.Writer,
) error {
	_, err := w.Write([]byte{0x3D, 0x30})

	return err
}
