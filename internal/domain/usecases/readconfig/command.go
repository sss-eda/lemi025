package readconfig

import (
	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
)

// CommandMessage TODO
type CommandMessage struct {
	InstrumentID  instrument.ID
	StationType   string
	StationNumber instrument.StationNumber
}

// Command TODO
func Command(
	instruments instrument.Repository,
) func(CommandMessage) error {
	return func(msg CommandMessage) error {
		instrument, err := instruments.Load(msg.InstrumentID)
		if err != nil {
			return err
		}
		event, err := instrument.Config.Read(
			msg.StationType,
			msg.StationNumber,
		)
		err = instruments.Save(msg.InstrumentID, event)
		if err != nil {
			return err
		}
		// log.Println(event)
		return nil
	}
}
