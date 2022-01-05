package gpio

import (
	"bufio"

	"github.com/sss-eda/lemi025"
)

// Driver TODO - analog for the TV in example. This is the concrete receiver
type Driver struct {
	*Connection
	state state
}

// NewDriver TODO
func NewDriver(
	connection *Connection,
) (*Driver, error) {
	return &Driver{
		connection,
		idle,
	}, nil
}

type state int

const (
	idle state = iota
	commandSent
	dataReceived
	configRead
	timeRead
	timeSet
)

// Drive TODO
func (driver *Driver) Drive(
	observer lemi025.Observer,
) {
	driver.RLock()
	defer driver.RUnlock()

	scanner := bufio.NewScanner(driver)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		b := scanner.Bytes()[0]

		switch driver.state {
		case idle:
			switch b {
			case 0x3F:
				driver.state = commandSent
			case 0x4C:
				driver.state = dataReceived
			}
		case commandSent:
			switch b {
			case 0x30:
				driver.state = configRead
			case 0x31:
				driver.state = timeRead
			case 0x32:
				driver.state = timeSet
			}
		case dataReceived:
		case configRead:

		}
	}
}
