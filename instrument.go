package lemi025

// Instrument TODO
type Instrument struct {
	readConfigCommand *driving.ReadConfigCommand
	readTimeCommand   *driving.ReadTimeCommand
	setTimeCommand    *driving.SetTimeCommand
}

// Drive TODO - Then in the driving repo/context we will have an interface:
// type Driver interface{
//     Drive()
// }
// type Instrument interface{
//     Driver
//     Observer
// }
func (instrument *Instrument) Drive() error {

}

// ReadConfig TODO
func (instrument *Instrument) ReadConfig() error {
	return instrument.readConfigCommand.Execute()
}

// ReadTime TODO
func (instrument *Instrument) ReadTime() error {
	return instrument.readTimeCommand.Execute()
}

// SetTime TODO
func (instrument *Instrument) SetTime() error {
	return instrument.setTimeCommand.Execute()
}

// type LEMI025 struct {
// 	Config        Config
// 	Time          Time
// 	Coefficients1 Coefficients1
// 	Coefficients2 Coefficients2
// 	GPSData       GPSData
// 	recording     bool
// 	FLASHData     FLASHData
// 	DAC           DAC
// }

// type Starter interface {
// 	Start() error
// }

// func (lemi *LEMI025) Start(
// 	adapter Starter,
// 	command StartCommand,
// ) error {
// 	if lemi.recording {
// 		return fmt.Errorf("unable to run \"Start\" command while recording")
// 	}

// 	return adapter.Start()
// }

// type Stopper interface {
// 	Stop() error
// }

// func (lemi *LEMI025) Stop(
// 	adapter Stopper,
// 	command StopCommand,
// ) error {
// 	err := adapter.Stop()
// 	if err != nil {
// 		return err
// 	}

// 	lemi.recording = false

// 	return nil
// }
