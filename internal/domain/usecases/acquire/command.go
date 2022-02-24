package acquire

import (
	"github.com/sss-eda/lemi025/internal/domain/entities/datum"
	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
	"github.com/sss-eda/lemi025/internal/domain/usecases/readconfig"
)

// CommandType TODO
type CommandType interface {
	readconfig.CommandMessage
}

// Command TODO
func Command[CT CommandType](
	instruments instrument.Repository,
	data datum.Repository,
) func(CT) error {
	return func(command CT) error {
		instrument, err := instruments.Load(msg.InstrumentID)
		if err != nil {
			return err
		}
		// event, err := instrument.
		err = instruments.Save(msg.InstrumentID, event)
		if err != nil {
			return err
		}
		// log.Println(event)
		return nil
	}
}
