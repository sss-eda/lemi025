package instrumentation

import (
	"sync"
)

// Time TODO
type Time struct {
	*sync.Mutex
	isInSync bool
	// events   []eventsourcing.Event
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
		Mutex:    &sync.Mutex{},
		isInSync: false,
		// events:   []eventsourcing.Event{},
		Year:   0,
		Month:  0,
		Day:    0,
		Hour:   0,
		Minute: 0,
		Second: 0,
	}

	return time
}

// OnSynchronised TODO
func (time *Time) OnSynchronised(event *TimeSynchronisedEvent) {
	time.Lock()
	defer time.Unlock()

	time.Year = event.Year
	time.Month = event.Month
	time.Day = event.Day
	time.Hour = event.Hour
	time.Minute = event.Minute
	time.Second = event.Second

	time.isInSync = true
}

// Read TODO
func (time *Time) Read() error {
	time.Lock()
	defer time.Unlock()

	time.isInSync = false

	return nil
}

// Set TODO
func (time *Time) Set() error {
	time.Lock()
	defer time.Unlock()

	time.isInSync = false

	return nil
}
