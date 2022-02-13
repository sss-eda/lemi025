package serial

import (
	"bufio"
	"io"
	"log"

	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
	"github.com/sss-eda/lemi025/internal/domain/usecases/updateconfig"
)

// Run TODO
func Run(
	r io.Reader,
	configReadEventer func(readconfig.Command) error,
	readTimeEventer func(updatetime.Command) error,
) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		data := scanner.Bytes()
		switch data[0] {
		case 0x4C: // "L" - Data frame received from instrument
			model := string(data[1:4])
			if model != "025" {
				log.Printf("expected LEMI-025 datapoint, but got LEMI-%q instead\n", model)
				continue
			}
		case 0x3F: // "?" - Response received from instrument
			switch data[1] {
			case 0x30: // "0" - Read Config
				err := configUpdater(updateconfig.Command{
					NewConfig: instrument.Config{
						StationNumber: instrument.StationNumber(data[2]),
					},
				})
				if err != nil {
					log.Println(err)
				}
			case 0x31: // "1" - Read Time
				err := configUpdater(updatetime.Command{
					NewConfig: instrument.Config{
						StationNumber: instrument.Time(data[2:8]),
					},
				})
				if err != nil {
					log.Println(err)
				}
			case 0x32: // "2" - Set Time
				err := configUpdater(updatetime.Command{
					NewConfig: instrument.Config{
						StationNumber: instrument.Time(data[2:8]),
					},
				})
				if err != nil {
					log.Println(err)
				}
			case 0x33: // "3" - Set Coefficients1
			}
		default:
			log.Printf("invalid input: %v\n", data)
			continue
		}
	}

	return scanner.Err()
}

// Handle TODO
func Synchronise(
	r io.Reader,
	// service synchronisation.Service,
) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		data := scanner.Bytes()
		switch data[0] {
		case 0x4C: // "L" - Data frame received from instrument
			model := string(data[1:4])
			if model != "025" {
				log.Printf("expected LEMI-025 datapoint, but got LEMI-%q instead\n", model)
				continue
			}
		case 0x3F: // "?" - Response received from instrument
			switch data[1] {
			case 0x30: // "0" - Read Config
				handler(updateconfig.Request{})
			case 0x31: // "1" - Read Time
				handler()
			case 0x32: // "2" - Set Time
			case 0x33: // "3" - Set Coefficients1
			}
		default:
			log.Printf("invalid input: %v\n", data)
			continue
		}
	}

}
