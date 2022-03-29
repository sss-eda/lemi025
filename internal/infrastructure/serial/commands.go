package serial

import (
	"context"
	"io"

	"github.com/sss-eda/lemi025"
)

// ReadConfig TODO
func ReadConfig(port io.Writer) func(context.Context, *lemi025.ReadConfigCommandPayload) error {
	return func(ctx context.Context, payload *lemi025.ReadConfigCommandPayload) error {
		_, err := port.Write([]byte{0x3D, 0x30})
		return err
	}
}

// ReadTime TODO
func ReadTime(port io.Writer) func(context.Context, *lemi025.ReadTimeCommandPayload) error {
	return func(ctx context.Context, payload *lemi025.ReadTimeCommandPayload) error {
		_, err := port.Write([]byte{0x3D, 0x31})
		return err
	}
}

// SetTime TODO
func SetTime(port io.Writer) func(context.Context, *lemi025.SetTimeCommandPayload) error {
	return func(ctx context.Context, payload *lemi025.SetTimeCommandPayload) error {
		_, err := port.Write([]byte{0x3D, 0x31})
		return err
	}
}
