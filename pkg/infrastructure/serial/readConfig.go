package serial

import (
	"io"

	"github.com/sss-eda/lemi025"
)

// ReadConfig TODO
func ReadConfig(
	w io.Writer,
) lemi025.ConfigReader {
	return func() error {
		return nil
	}
}
