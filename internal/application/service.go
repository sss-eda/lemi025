package driving

import (
	"context"

	"github.com/sss-eda/lemi025/internal/domain/driving"
)

// Driver TODO
type Driver struct {
	Instruments driving.InstrumentRepository
}

// ReadConfig TODO
func (driver *Driver) ReadConfig(
	ctx context.Context,
	sn driving.StationNumber,
) error {
	instrument, err := driver.Instruments.GetInstrumentByStationNumber(ctx, sn)
	if err != nil {
		return err
	}

	instrument.Write()

	return nil
}

// ReadTime TODO
func (driver *Driver) ReadTime() error {
	return nil
}
