package serial

import (
	"context"
	"io"

	"github.com/sss-eda/lemi025"
)

// ReadConfigCommandHandler TODO
func ReadConfigCommandHandler(port io.Writer) func(context.Context, *lemi025.ReadConfigCommandPayload) error {
	return func(ctx context.Context, payload *lemi025.ReadConfigCommandPayload) error {
		buffer := make([]byte, 2)

		buffer[0] = sendToken
		buffer[1] = readConfigToken

		_, err := port.Write(buffer)

		return err
	}
}

// ReadTime TODO
func ReadTime(port io.Writer) func(context.Context, *lemi025.ReadTimeCommandPayload) error {
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
func SetTime(port io.Writer) func(context.Context, *lemi025.SetTimeCommandPayload) error {
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

// SetCoefficients1 TODO
func SetCoefficients1(port io.Writer) func(context.Context, *lemi025.SetCoefficients1CommandPayload) error {
	return func(
		ctx context.Context,
		payload *lemi025.SetCoefficients1CommandPayload,
	) error {
		buffer := make([]byte, 4)

		buffer[0] = sendToken
		buffer[1] = setCoefficients1Token
		buffer[2] = 0x00
		buffer[3] = byte(payload.Mode)

		_, err := port.Write(buffer)

		return err
	}
}

// ReadCoefficients1 TODO
func ReadCoefficients1(port io.Writer) func(context.Context, *lemi025.ReadCoefficients1CommandPayload) error {
	return func(
		ctx context.Context,
		payload *lemi025.ReadCoefficients1CommandPayload,
	) error {
		buffer := make([]byte, 6)

		buffer[0] = sendToken
		buffer[1] = readCoefficients1Token
		buffer[2] = 0x00
		buffer[3] = byte(payload.Mode)
		buffer[4] = byte(payload.Uin)
		buffer[5] = byte(payload.Mode1)

		_, err := port.Write(buffer)

		return err
	}
}
