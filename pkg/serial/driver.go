package serial

import (
	"context"
	"io"

	"github.com/sss-eda/lemi025"
	tarm "github.com/tarm/serial"
)

type Driver struct {
	io.ReadWriteCloser
	onDataFrameAcquired []func(lemi025.DataFrameAcquiredEvent)
	onConfigRead        []func(lemi025.ConfigReadEvent)
	onTimeRead          []func(lemi025.TimeReadEvent)
	onTimeSet           []func(lemi025.TimeSetEvent)
	onCoefficients1Set  []func(lemi025.Coefficients1SetEvent)
	onCoefficients1Read []func(lemi025.Coefficients1ReadEvent)
	onCoefficients2Set  []func(lemi025.Coefficients2SetEvent)
	onCoefficients2Read []func(lemi025.Coefficients2ReadEvent)
	onGPSDataRead       []func(lemi025.GPSDataReadEvent)
	onStarted           []func(lemi025.StartedEvent)
	onFLASHChecked      []func(lemi025.FLASHCheckedEvent)
}

func NewDriver(
	ctx context.Context,
	config Config,
) (*Driver, error) {
	port, err := tarm.OpenPort(&tarm.Config{
		Name: config.Name,
		Baud: config.Baud,
	})
	if err != nil {
		return nil, err
	}

	driver := Driver{
		ReadWriteCloser:     port,
		onDataFrameAcquired: []func(lemi025.DataFrameAcquiredEvent){},
		onConfigRead:        []func(lemi025.ConfigReadEvent){},
		onTimeRead:          []func(lemi025.TimeReadEvent){},
		onTimeSet:           []func(lemi025.TimeSetEvent){},
		onCoefficients1Set:  []func(lemi025.Coefficients1SetEvent){},
		onCoefficients1Read: []func(lemi025.Coefficients1ReadEvent){},
		onCoefficients2Set:  []func(lemi025.Coefficients2SetEvent){},
		onCoefficients2Read: []func(lemi025.Coefficients2ReadEvent){},
		onGPSDataRead:       []func(lemi025.GPSDataReadEvent){},
		onStarted:           []func(lemi025.StartedEvent){},
		onFLASHChecked:      []func(lemi025.FLASHCheckedEvent){},
	}
	go driver.loop(ctx)

	return &driver, nil
}

func (driver *Driver) loop(
	ctx context.Context,
) error {
	return nil
}

func (driver *Driver) ReadConfig(
	command lemi025.ReadConfigCommand,
) error {
	return nil
}

func (driver *Driver) OnConfigRead(
	callback func(lemi025.ConfigReadEvent),
) {
	driver.onConfigRead = append(driver.onConfigRead, callback)
}

func (driver *Driver) ReadTime(
	command lemi025.ReadConfigCommand,
) error {
	return nil
}
func (driver *Driver) OnTimeRead(
	callback func(lemi025.TimeReadEvent),
) {
	driver.onTimeRead = append(driver.onTimeRead, callback)
}

func (driver *Driver) SetTime(
	command lemi025.SetTimeCommand,
) error {
	return nil
}

func (driver *Driver) OnTimeSet(
	callback func(lemi025.TimeSetEvent),
) {
	driver.onTimeSet = append(driver.onTimeSet, callback)
}
