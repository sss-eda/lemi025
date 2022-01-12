package driving

import (
	"context"
	"log"
)

// Service TODO
type Service interface {
	Drive(context.Context, DeviceRepository)
}

// DeviceRepository TODO
type DeviceRepository interface {
	Load(DeviceID) (*Device, error)
	Save(DeviceID, *Device) error
}

type service struct {
	devices DeviceRepository
}

// Drive TODO
func (driver *service) Drive(
	ctx context.Context,
	id DeviceID,
) error {
	device, err := driver.devices.Load(id)
	if err != nil {
		return err
	}

	device.time.TimeUpdatedEvent.Register(
		func(payload TimeUpdatedEventPayload) {
			log.Println(payload)
		},
	)

	return nil
}
