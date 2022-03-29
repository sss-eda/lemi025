package serial

import (
	"bufio"
	"context"
	"io"

	"github.com/sss-eda/lemi025"
)

// Listen TODO
func Listen(
	onConfigRead func(context.Context, *lemi025.ConfigReadEventPayload) error,
	onTimeRead func(context.Context, *lemi025.TimeReadEventPayload) error,
	onTimeSet func(context.Context, *lemi025.TimeSetEventPayload) error,
) func(context.Context, io.Reader) error {
	return func(ctx context.Context, port io.Reader) error {
		scanner := bufio.NewScanner(port)
		scanner.Split(bufio.ScanBytes)

		for scanner.Scan() {
			switch scanner.Bytes()[0] {
			case 0x3F:
			case 0x4C:
			default:
			}
		}

		return scanner.Err()
	}
}
