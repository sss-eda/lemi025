package serial

import (
	"io"
)

// Station TODO
type Station struct {
	io.Writer
	io.Reader
}

// // NewStationDriver TODO
// func NewStation(reader io.Reader, writer io.Writer) (*StationDriver, error) {
// 	return &StationDriver{
// 		writer,
// 		reader,
// 	}, nil
// }

// ReadConfig TODO
func (station *Station) ReadConfig() error {
	_, err := station.Write([]byte{0x3D, 0x30})
	if err != nil {
		return err
	}

	return nil
}
