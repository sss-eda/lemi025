package text

// File TODO
type File struct {
	name string
	data []*DataPoint
}

// NewFileFromEvents TODO
func NewFileFromEvents(
	eventstreams <-chan <-chan Event,
) {
	for events := range eventstreams {
		for event := range events {

		}
	}
}

// OnNewDataPointAcquired TODO
func (file *File) OnNewDataPointAcquired(
	event acquiring.NewDataPointAcquiredEvent,
) {

}
