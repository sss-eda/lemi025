package gpio

import (
	"bufio"
	"context"

	"github.com/sss-eda/lemi025/driving"
)

// Driver TODO
type Driver struct {
	device     *driving.Device
	connection Connection
	state      state
}

type state interface {
	next(byte) state
}

// NewDriver TODO
func NewDriver(
	ctx context.Context,
	device *driving.Device,
	connection Connection,
	events chan<- Event,
) (*Driver, error) {
	driver := &Driver{
		device:     device,
		connection: connection,
	}

	go driver.loop(ctx, events)

	return driver, nil
}

func (driver *Driver) loop(
	ctx context.Context,
) {
	scanner := bufio.NewScanner(driver.connection)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		b := scanner.Bytes()[0]
		switch driver.state {
		case idle:

		}
	}
}

// ReadConfig TODO
func (driver *Driver) ReadConfig() error {
	_, err := driver.connection.Write([]byte{0x3D, 0x30})
	if err != nil {
		return err
	}

	return nil
}

// ReadTime TODO
func (driver *Driver) ReadTime() error {
	_, err := driver.connection.Write([]byte{0x3D, 0x31})
	if err != nil {
		return err
	}

	return nil
}

// SetTime TODO
func (driver *Driver) SetTime(
	time driving.Time,
) error {
	_, err := driver.connection.Write([]byte{0x3D, 0x32})
	if err != nil {
		return err
	}

	return nil
}
