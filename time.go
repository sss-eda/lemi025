package lemi025

// Time TODO - This is a value object.
//   * Having the setter and reader only makes sense in the
//     "DRIVING"/"CONTROLLING"/"COMMANDING" context
type Time struct {
	reader Command
	setter Command
	sync   bool
	Year   uint16 `json:"year"`
	Month  uint8  `json:"month"`
	Day    uint8  `json:"day"`
	Hour   uint8  `json:"hour"`
	Minute uint8  `json:"minute"`
	Second uint8  `json:"second"`
}

// NewTime TODO
func NewTime(
	timeReader Command,
	timeSetter Command,
) *Time {
	time := &Time{
		reader: timeReader,
		setter: timeSetter,
		sync:   true,
	}

	// go func() {
	// 	for event := range events {
	// 		switch e := event.(type) {
	// 		case TimeReadEvent:
	// 			time.Year = e.year
	// 			time.Month = e.month
	// 		case TimeSetEvent:
	// 			time.Year = e.year
	// 			time.Month = e.month
	// 		}
	// 	}
	// }()

	return time
}

// Read TODO
func (time *Time) Read() error {
	err := time.reader.Execute()
	if err != nil {
		return err
	}

	time.sync = false

	return nil
}

// Set TODO
func (time *Time) Set() error {
	// The set-time command only works when device is not recording. Return an
	// error in such a case.
	// if device.IsRecording() {
	// 	return fmt.Errorf("cannot set time while device is recording")
	// }

	err := time.setter.Execute()
	if err != nil {
		return err
	}

	// If the command was sent successfully, set the sync state of the device's
	// time to false. When the device hears back from the lemi, it will be
	// returned to a synchronised state.
	time.sync = false

	return nil
}
