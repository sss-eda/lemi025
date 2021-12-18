package lemi025

type Driver interface {
	ReadConfig(ReadConfigCommand) error
	ReadTime(ReadTimeCommand) error
	SetTime(SetTimeCommand) error
	SetCoefficients1(SetCoefficients1Command) error
	ReadCoefficients1(ReadCoefficients1Command) error
	SetCoefficients2(SetCoefficients2Command) error
	ReadCoefficients2(ReadCoefficients2Command) error
	ReadGPSData(ReadGPSDataCommand) error
	Stop(StopCommand) error
	Start(StartCommand) error
	CheckFLASH(CheckFLASHCommand) error
	SetDACx(SetDACxCommand) error
	SetDACy(SetDACyCommand) error
	SetDACz(SetDACzCommand) error
	OnDataFrameAcquired(func(DataFrameAcquiredEvent))
	OnConfigRead(func(ConfigReadEvent))
	OnTimeRead(func(TimeReadEvent))
	OnTimeSet(func(TimeSetEvent))
	OnCoefficients1Set(func(Coefficients1SetEvent))
	OnCoefficients1Read(func(Coefficients1ReadEvent))
	OnCoefficients2Set(func(Coefficients2SetEvent))
	OnCoefficients2Read(func(Coefficients2ReadEvent))
	OnGPSDataRead(func(GPSDataReadEvent))
	OnStarted(func(StartedEvent))
	OnFLASHChecked(func(FLASHCheckedEvent))
}
