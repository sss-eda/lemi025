package eventsourcing

import "github.com/sss-eda/lemi025/internal/domain"

// ReadConfig TODO
func ReadConfig(
	store Store,
) func(domain.ReadConfigInput) error {
	return func(input domain.ReadConfigInput) error {
		instrument, err := store.Load(input.ID)
		if err != nil {
			return err
		}

		err = instrument.Config.Read(strategy(input))
		if err != nil {
			return err
		}
	}
}
