package serial

import (
	"bufio"
	"io"

	"github.com/sss-eda/lemi025"
)

// type Station struct {
// 	responses [][]byte
// 	buffer    []byte
// }

// func NewStation(port io.ReadWriteCloser) (*Station, error) {
// 	station := Station{
// 		responses: [][]byte{},
// 		buffer:    []byte{},
// 	}

// 	// go func() {
// 	// 	p := []byte{}
// 	// 	var n int
// 	// 	var err error
// 	// 	for {
// 	// 		_, err = port.Read(p)
// 	// 		if err != nil {
// 	// 			log.Printf("failed to read from serial port: %v\n", err)
// 	// 			continue
// 	// 		}
// 	// 		for
// 	// 	}
// 	// }()

// 	return &station, nil
// }

// func (station *Station) ReadConfig(command lemi025.ReadConfigCommand) error {
// 	_, err := station.writer.Write([]byte{0x3D, 0x30})
// 	return err
// }

func Lemi025(port io.ReadWriteCloser) (<-chan lemi025.Event, chan<- lemi025.Command) {
	commands := make(chan lemi025.Command)
	events := make(chan lemi025.Event)

	go func() {
		for command := range commands {
			switch command.Kind {
			case lemi025.ReadConfigCommand:
				readConfig(port)
			case lemi025.ReadTimeCommand:
				readTime(port)
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(port)
		scanner.Split(bufio.ScanBytes)

		for scanner.Scan() {
			b := scanner.Bytes()[0]
		}
	}()
}
