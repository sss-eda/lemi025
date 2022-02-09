package serial

import (
	"bufio"
	"fmt"
	"io"

	"github.com/sss-eda/lemi025/internal/application"
)

// Transmitter TODO
func Transmitter(
	w io.Writer,
) func(application.Command) error {
	return func[T application.CommandType](command T) error {
		switch command.Type() {
		case application.ReadConfig:
			_, err := w.Write([]byte{0x3D, 0x30})
			if err != nil {
				return err
			}
		case application.ReadTime:
			_, err := w.Write([]byte{0x3D, 0x31})
			if err != nil {
				return err
			}
		case application.SetTime:
			data := command.Payload()
			
			_, err := w.Write([)
		default:
			return fmt.Errorf("invalid command type")
		}

		return nil
	}
}

// Receiver TODO
func Receiver(
	r io.Reader,
) func(*domain.Instrument) error {
	return func(instrument *domain.Instrument) error {
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			b := scanner.Bytes()[0]
			//...
			instrument.Config.Sync(b)
		}
	}
}