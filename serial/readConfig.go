package serial

import "github.com/sss-eda/lemi025"

func ReadConfig(command lemi025.ReadConfigCommand) error {

}

func OnConfigRead(station lemi025.Station, dispatch func(lemi025.ConfigReadEvent)) error {
	
}

station.OnConfigRead(jetstream.Dispatch("instrumentation.lemi025.1.events"))