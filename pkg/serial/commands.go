package serial

import (
	"fmt"
	"io"

	"github.com/sss-eda/lemi025"
)

type Transmitter interface {
	Transmit(lemi025.Command) error
}

type AsyncTransmitter interface {
	TransmitAsync(
		done <-chan interface{},
		commands <-chan lemi025.Command,
	) <-chan error
}

type Receiver interface {
	Receive() lemi025.Event
}

type AsyncReceiver interface {
	ReceiveAsync(
		done <-chan interface{},
		events chan<- lemi025.Event,
	)
}

type transmitter struct {
	io.Writer
}

func (tx *transmitter) T(
	command lemi025.Command,
) error {

}

func (ch *commandHandler) HandleCommands(
	done <-chan interface{},
	commands <-chan lemi025.Command,
) chan<- error {
	errors := make(chan error)

	go func() {
		defer close(eventStream)
		var err error

		for command := range commands {
			switch command := command.(type) {
			case lemi025.ReadConfigCommand:
				err = m.ReadConfig(command)
			case lemi025.ReadTimeCommand:
				err = m.ReadTime(command)
			case lemi025.SetTimeCommand:
				err = m.SetTime()
			default:
				err = fmt.Errorf("unsupported command: %v", command)
			}

			select {
			case <-done:
				return
			case errors <- err:
			}
		}
	}()

	return errors
}

func (ch *commandHandler) GetEvents(
	done <-chan interface{},
	events chan<- lemi025.Event,
) {
	go func() {
		scanner := bufio.NewScanner()
		for scanner.Scan() {

		}
	}
}
