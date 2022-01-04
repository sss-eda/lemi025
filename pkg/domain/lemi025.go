package lemi025

import "time"

// LEMI025 TODO
type LEMI025 interface {
	ReadConfig() error
	ReadTime() error
	SetTime(time.Time) error
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
