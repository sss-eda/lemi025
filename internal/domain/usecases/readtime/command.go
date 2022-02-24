package readconfig

import (
	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
)

// CommandMessage TODO
type CommandMessage struct {
	InstrumentID instrument.ID
	Year         instrument.Year
	Month        instrument.Month
	Day          instrument.Day
	Hour         instrument.Hour
	Minute       instrument.Minute
	Second       instrument.Second
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
		event, err := instrument.Time.Read(
			msg.Year,
			msg.Month,
			msg.day,
			msg.Hour,
			msg.Minute,
			msg.Second,
		)
		err = instruments.Save(msg.InstrumentID, event)
		if err != nil {
			return err
		}
		// log.Println(event)
		return nil
	}
}
