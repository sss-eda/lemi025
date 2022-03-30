package serial

import (
	"context"
	"io"

	"github.com/sss-eda/lemi025"
)

// ReadConfig TODO
func ReadConfig(
	port io.Writer,
) func(context.Context, *lemi025.ReadConfigCommandPayload) error {
	return func(
		ctx context.Context,
		payload *lemi025.ReadConfigCommandPayload,
	) error {
		buffer := make([]byte, 2)

		buffer[0] = sendToken
		buffer[1] = readConfigToken

		_, err := port.Write(buffer)

		return err
	}
}

// ReadTime TODO
func ReadTime(
	port io.Writer,
) func(context.Context, *lemi025.ReadTimeCommandPayload) error {
	return func(
		ctx context.Context,
		payload *lemi025.ReadTimeCommandPayload,
	) error {
		buffer := make([]byte, 2)

		buffer[0] = sendToken
		buffer[1] = readTimeToken

		_, err := port.Write(buffer)

		return err
	}
}

// SetTime TODO
func SetTime(
	port io.Writer,
) func(context.Context, *lemi025.SetTimeCommandPayload) error {
	return func(
		ctx context.Context,
		payload *lemi025.SetTimeCommandPayload,
	) error {
		buffer := make([]byte, 8)

		buffer[0] = sendToken
		buffer[1] = setTimeToken
		buffer[2] = payload.Year
		buffer[3] = payload.Month
		buffer[4] = payload.Day
		buffer[5] = payload.Hour
		buffer[6] = payload.Minute
		buffer[7] = payload.Second

		_, err := port.Write(buffer)

		return err
	}
}

func 