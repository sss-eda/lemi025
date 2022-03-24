package serial

import (
	"bufio"
	"context"
	"io"

	"github.com/sss-eda/lemi025/internal/receiving"
)

// Receive TODO
func Receive(
	ctx context.Context,
	station *receiving.Station,
	port io.Reader,
) {
	scanner := bufio.NewScanner(port)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		b := scanner.Bytes()[0]

		receiving.ConfigReadEvent{
			StationNumber: b,
		}.Apply(ctx, station.Config)
	}
}
