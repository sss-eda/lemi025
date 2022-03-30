package serial

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/sss-eda/lemi025"
)

// Subscribe TODO
func Subscribe(
	ctx context.Context,
	port io.Reader,
	onDatumAcquired func(context.Context, *lemi025.DatumAcquiredEventPayload),
	onConfigRead func(context.Context, *lemi025.ConfigReadEventPayload),
	onTimeRead func(context.Context, *lemi025.TimeReadEventPayload),
	onTimeSet func(context.Context, *lemi025.TimeSetEventPayload),
) error {
	scanner := bufio.NewScanner(port)
	scanner.Split(bufio.ScanBytes)

	buffer := []byte{}
	state := idleToken
	for scanner.Scan() {
		token := scanner.Bytes()[0]
		select {
		case <-ctx.Done():
			break
		default:
			switch state {
			case idleToken:
				switch token {
				case receiveToken:
					state = receiveToken
				case acquireToken:
					state = acquireToken
				default:
					log.Printf("unexpected token: %v\n", token)
					continue
				}
			case receiveToken:
				switch token {
				case readConfigToken:
					state = readConfigToken
				case readTimeToken:
					state = readTimeToken
				case setTimeToken:
					state = setTimeToken
				default:
					state = idleToken
				}
			case acquireToken:
				state = acquireToken
			case readConfigToken:
				buffer = append(buffer, token)
				if len(buffer) < 6 {
					buffer = append(buffer, token)
				}
			case readTimeToken:
				if len(buffer) < 6 {
					buffer = append(buffer, token)
					continue
				}
				event, err := newTimeReadEventPayload(buffer)
			}
		}
	}

	return scanner.Err()
}

func newConfigReadEventPayload(buffer []byte) (*lemi025.ConfigReadEventPayload, error) {
	stationType := string(buffer[:4])
	stationNumber := uint8(buffer[4])

	if stationType != "L025" {
		return nil, fmt.Errorf("unexpected station type: %v", stationType)
	}

	return &lemi025.ConfigReadEventPayload{
		StationType:   stationType,
		StationNumber: stationNumber,
	}, nil
}
