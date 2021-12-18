package lemi025

type Time struct {
	year   uint16
	month  uint8
	day    uint8
	hour   uint8
	minute uint8
	second uint8
}

type TimeReader interface {
	ReadTime() error
}

func (time *Time) Read(
	reader TimeReader,
	command ReadTimeCommand,
) error {
	return reader.ReadTime()
}

type TimeSetter interface {
	SetTime(year uint16, month, day, hour, minute, second uint8) error
}

func (time *Time) Set(
	setter TimeSetter,
	command SetTimeCommand,
) error {
	return setter.SetTime(
		command.Year,
		command.Month,
		command.Day,
		command.Hour,
		command.Minute,
		command.Second,
	)
}
