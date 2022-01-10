package lemi025

// Time TODO - This is a value object.
//   * Having the setter and reader only makes sense in the
//     "DRIVING"/"CONTROLLING"/"COMMANDING" context
type Time struct {
	sync   bool
	Year   uint16 `json:"year"`
	Month  uint8  `json:"month"`
	Day    uint8  `json:"day"`
	Hour   uint8  `json:"hour"`
	Minute uint8  `json:"minute"`
	Second uint8  `json:"second"`
}

// NewTime TODO
func NewTime() *Time {
	time := &Time{
		sync: false,
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

// Sync TODO
func (time *Time) Sync() error {
	time.sync = true

	return nil
}

// Read TODO
func (time *Time) Desync() error {
	time.sync = false

	return nil
}
