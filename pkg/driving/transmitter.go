package serial

import (
	"io"

	"github.com/sss-eda/lemi025"
)

type transmitter struct {
	io.Writer
}

func (tx *transmitter) Transmit(
	input lemi025.Command,
) error {
	switch command := input.(type) {
	case lemi025.ReadConfigCommand:
		readConfig(tx)
	case lemi025.ReadTimeCommand:
		readConfig(tx)
	case lemi025.SetTimeCommand:
	}
}

func (tx *transmitter) TransmitAsync(
	done <-chan interface{},
	commands <-chan lemi025.Command,
) {

}
