package driving

import "github.com/sss-eda/lemi025"

// Service TODO
type Service struct {
	readConfigCommand lemi025.Command
	readTimeCommand   lemi025.Command
	setTimeCommand    lemi025.Command
}

// CommandFactory TODO
type CommandFactory interface {
	NewReadConfigCommand(*Device) lemi025.Command
	NewReadTimeCommand(*Device) lemi025.Command
	NewSetTimeCommand(*Device) lemi025.Command
}

// NewService TODO
func NewService(
	commandFactory CommandFactory,
	device *Device,
) *Service {
	service := &Service{
		readConfigCommand: commandFactory.NewReadConfigCommand(device),
		readTimeCommand:   commandFactory.NewReadTimeCommand(device),
		setTimeCommand:    commandFactory.NewSetTimeCommand(device),
	}

	return service
}

// ReadConfig TODO
func (service *Service) ReadConfig(
	device *Device,
) error {
	return service.readConfigCommand.Execute()
}

// ReadTime TODO
func (service *Service) ReadTime(
	device *Device,
) error {
	return service.readTimeCommand.Execute()
}

// SetTime TODO
func (service *Service) SetTime(
	device *Device,
) error {
	return service.setTimeCommand.Execute()
}
