package readtime

import (
	"context"

	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
)

// Command TODO
type Command struct {
	InstrumentID instrument.ID
	Year         instrument.Year
	Month        instrument.Month
	Day          instrument.Day
	Hour         instrument.Hour
	Minute       instrument.Minute
	Second       instrument.Second
}

// Mutate TODO
func Mutate(
	load func(context.Context, instrument.ID) (*instrument.Instrument, error),
	save func(context.Context, instrument.ID, instrument.Event) error,
) func(Command) error {
	return func(command Command) error {
		_, err := load(command.InstrumentID)
		if err != nil {
			return err
		}
		// event, err := instrument.Time.Read(
		// 	command.Year,
		// 	command.Month,
		// 	command.day,
		// 	command.Hour,
		// 	command.Minute,
		// 	command.Second,
		// )
		// err = instruments.Save(command.InstrumentID, event)
		// if err != nil {
		// 	return err
		// }
		// log.Println(event)
		return nil
	}
}
