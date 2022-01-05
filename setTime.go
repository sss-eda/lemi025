package lemi025

// TimeSetter TODO
type TimeSetter interface {
	SetTime(
		year uint16,
		month uint8,
		day uint8,
		hour uint8,
		minute uint8,
		second uint8,
	) error
}

// SetTimeCommand TODO
type SetTimeCommand struct {
	TimeSetter `json:"-"`
	Year       uint16 `json:"year"`
	Month      uint8  `json:"month"`
	Day        uint8  `json:"day"`
	Hour       uint8  `json:"hour"`
	Minute     uint8  `json:"minute"`
	Second     uint8  `json:"second"`
}

// Execute TODO
func (command *SetTimeCommand) Execute() error {
	return command.SetTime(
		command.Year,
		command.Month,
		command.Day,
		command.Hour,
		command.Minute,
		command.Second,
	)
}
